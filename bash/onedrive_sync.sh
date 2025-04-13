#!/bin/bash
while true; do
    echo "=== Sync started at $(date) ===" >> ~/onedrive_sync.log
    rclone bisync ~/Documents/Obsidian onedrive:Obsidian \
      --conflict-resolve path1 \
      --conflict-loser delete \
      --backup-dir1 ~/Documents/Obsidian_deleted_local \
      --backup-dir2 onedrive:Obsidian_deleted_remote \
      --exclude "*.swp" --exclude "*.tmp" --exclude ".DS_Store" --exclude ".~lock.*" \
      --progress >> ~/onedrive_sync.log 2>&1


	code=$?

    if [[ $code -eq 3 ]]; then
        echo "=== CRITICAL ERROR, calling RESYNC SCRIPT at $(date) ===" >> ~/onedrive_sync.log
        bash /home/davidapdf/projects/bash/onedrive_cash.sh >> ~/onedrive_sync.log 2>&1
    fi
    sleep 30

done