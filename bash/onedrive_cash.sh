rclone bisync ~/Documents/Obsidian onedrive:Obsidian \
  --conflict-resolve path1 \
  --conflict-loser pathname \
  --backup-dir1 ~/Documents/Obsidian_deleted_local \
  --backup-dir2 onedrive:Obsidian_deleted_remote \
  --exclude "*.swp" --exclude "*.tmp" --exclude ".DS_Store" --exclude ".~lock.*" \
  --resync \
  --progress


