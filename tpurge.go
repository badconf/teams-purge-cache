package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check if Microsoft Teams is running
	isRunning, err := isTeamsRunning()
	if err != nil {
		fmt.Println("Error checking if Teams is running:", err)
		return
	}

	if isRunning {
		// Inform the user to close Microsoft Teams
		fmt.Println("Microsoft Teams is running. Please close it and run the program again.")
		return
	}

	// Get the AppData path
	appDataPath := os.Getenv("APPDATA")

	// Folders to check
	folders := []string{
		appDataPath + "\\Microsoft\\teams\\tmp",
		appDataPath + "\\Microsoft\\teams\\Blob_storage",
		appDataPath + "\\Microsoft\\teams\\Cache",
		appDataPath + "\\Microsoft\\teams\\IndexedDB",
		appDataPath + "\\Microsoft\\teams\\GPUCache",
		appDataPath + "\\Microsoft\\teams\\databases",
		appDataPath + "\\Microsoft\\teams\\Local Storage",
	}

	for _, folder := range folders {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Skip the root folder itself and any subdirectories
			if path == folder || info.IsDir() {
				return nil
			}
			// Delete the file
			err = os.Remove(path)
			if err != nil {
				return err
			}
			fmt.Println("Deleted file:", path)
			return nil
		})
		if err != nil {
			fmt.Println("Error processing folder", folder, ":", err)
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
	return strings.Contains(string(out), "Teams.exe"), nil
}
