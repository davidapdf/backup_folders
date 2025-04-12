package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)


func runCommand(name string, args ...string) error {
    configPath := "/home/davidapdf/.config/rclone/rclone.conf"
    
    // Prepend the --config flag to the rclone arguments
    finalArgs := append([]string{"--config", configPath}, args...)

    cmd := exec.Command(name, finalArgs...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func main() {
	//CONFIG
	localBase := "/home/davidapdf/bkp_tmp/onedrive_snapshots"
	onedriveFolders := []struct {
		remotePath string
		localName string
	}{
		{"temp","temp"},
	}
	googleDriveRemote := "gdrive:OneDrive_Snapshots"

	//Snapshot Folder
	date := time.Now().Format("2006-01-02")
	snapshotFolder := filepath.Join(localBase, "onedrive_snapshot_"+date)
	if err := os.MkdirAll(snapshotFolder, 0775); err != nil {
		log.Fatalf("Failed to create snapshot folder: %v", err)
	}

	//Step 1: Download selected folders from OneDrive
	for _, folder := range onedriveFolders {
		localPath := filepath.Join(snapshotFolder, folder.localName)
		remote := fmt.Sprintf("onedrive:%s",folder.remotePath)
		fmt.Printf("Copying from %s to %s\n", remote, localPath)
		err := runCommand("rclone", "copy", remote, localPath, "--progress")
		if err != nil {
			log.Fatalf("Failed to copy folder: %v", err)
		}

	}
	// Step 2: Upload snapshot to Google Drive
	fmt.Printf("Uploading snapshot to Google Drive...\n")
	err := runCommand("rclone", "copy", snapshotFolder, fmt.Sprintf("%s/onedrive_snapshot_%s", googleDriveRemote, date), "--progress")
	if err != nil {
		log.Fatalf("Failed to upload snapshot to Google Drive: %v", err)
	}
	// Step 3: Optional clean up
	fmt.Println("Cleaning up local snapshot...")
	err = os.RemoveAll(snapshotFolder)
	if err != nil {
		log.Printf("Warning: failed to remove local snapshot: %v", err)
	}

	fmt.Println("Backup complete!")
}