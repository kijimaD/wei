package wei

import (
	"os"
	"path/filepath"
	"strings"
)

const homeDir = "~"

func ExpandHomedir(original string) string {
	expanded := original
	if strings.HasPrefix(original, homeDir) {
		dirname, _ := os.UserHomeDir()
		expanded = filepath.Join(dirname, original[len(homeDir):])
	}
	return expanded
}
