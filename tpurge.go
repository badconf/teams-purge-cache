package main

import (
	"fmt"
	"os"
	"os/exec"
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

	// Teams cache folder
	teamsCacheFolder := appDataPath + "\\Microsoft\\Teams"

	// Delete the Teams cache folder
	err = os.RemoveAll(teamsCacheFolder)
	if err != nil {
		fmt.Println("Error deleting Teams cache folder:", err)
		return
	}

	fmt.Println("Teams cache folder deleted successfully.")
}

// isTeamsRunning checks if Microsoft Teams is running
func isTeamsRunning() (bool, error) {
	out, err := exec.Command("tasklist", "/FI", "IMAGENAME eq Teams.exe").Output()
	if err != nil {
		return false, err
	}
	return strings.Contains(string(out), "Teams.exe"), nil
}
