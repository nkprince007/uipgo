package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

// GetMacOSWallpaper returns the path to the current wallpaper.
func GetMacOSWallpaper() (string, error) {
	// notest

	stdout, err := exec.Command(
		"osascript", "-e",
		`tell application "Finder" to get POSIX path of (get desktop picture as alias)`,
	).Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

// SetMacOSWallpaper sets the wallpaper to the picture from provided path.
func SetMacOSWallpaper(path string) error {
	// notest

	return exec.Command(
		"osascript", "-e",
		`tell application "Finder" to set desktop picture to POSIX file `+
			strconv.Quote(path)).Run()
}
