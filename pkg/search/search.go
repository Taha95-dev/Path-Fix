package search

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Common install paths for different OS
var commonPaths = map[string][]string{
	"windows": {
		os.Getenv("USERPROFILE") + `\scoop\apps`,
		os.Getenv("USERPROFILE") + `\go\bin`,
		`C:\Program Files`,
		`C:\Program Files (x86)`,
		os.Getenv("LOCALAPPDATA") + `\Programs`,
	},
	"linux": {
		"/usr/bin",
		"/usr/local/bin",
		"/opt",
		os.Getenv("HOME") + "/.local/bin",
		os.Getenv("HOME") + "/go/bin",
	},
	"darwin": { // macOS
		"/usr/local/bin",
		"/opt/homebrew/bin",
		os.Getenv("HOME") + "/go/bin",
	},
}

func InPATH(executable string) bool {
	os := runtime.GOOS

	switch os {
	case "windows":
		winPath := executable + ".exe"
		_, err := exec.LookPath(winPath)
		return err == nil
	default:
	_, err := exec.LookPath(executable)
	return err == nil
	}
}
func FindInPaths(executable string) string {
	OS := runtime.GOOS

	if InPATH(executable) {
	return ""
	}
	paths := commonPaths[OS]
	for _, basePath := range paths {
		fullPath := filepath.Join(basePath, executable)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath
		}

		if runtime.GOOS == "windows" {
			fullPathExe := filepath.Join(basePath, executable + ".exe")
			if _,err := os.Stat(fullPathExe); err == nil {
				return fullPathExe
			}
		}
	}
	return ""
}








