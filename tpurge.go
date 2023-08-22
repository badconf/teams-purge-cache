package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Check if Microsoft Teams is running
	isRunning, err := isTeamsRunning()
	if err != nil {
		fmt.Println("Error checking if Teams is running:", err)
		return
	}

	if isRunning {
		// Close Microsoft Teams
		err := exec.Command("taskkill", "/F", "/IM", "Teams.exe").Run()
		if err != nil {
			fmt.Println("Failed to close Teams:", err)
			return
		}
	}

	// Folders to purge
	folders := []string{
		"%appdata%\\Microsoft\\teams\\tmp",
		"%appdata%\\Microsoft\\teams\\Blob_storage",
		"%appdata%\\Microsoft\\teams\\Cache",
		"%appdata%\\Microsoft\\teams\\IndexedDB",
		"%appdata%\\Microsoft\\teams\\GPUCache",
		"%appdata%\\Microsoft\\teams\\databases",
		"%appdata%\\Microsoft\\teams\\Local Storage",
		"%appdata%\\Microsoft\\teams\\Application Cache\\Cache",
	}

	for _, folder := range folders {
		// Expand environment variables
		folderPath := os.ExpandEnv(folder)
		err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Delete files and folders
			return os.RemoveAll(path)
		})
		if err != nil {
			fmt.Println("Error deleting files in folder", folder, ":", err)
			return
		}
	}

	fmt.Println("Operation completed successfully.")
}

// isTeamsRunning checks if Microsoft Teams is running
func isTeamsRunning() (bool, error) {
	out, err := exec.Command("tasklist", "/FI", "IMAGENAME eq Teams.exe").Output()
	if err != nil {
		return false, err
	}
	return len(out) > 0, nil
}
