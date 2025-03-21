package config

import (
	"os"
	"path/filepath"
)

const directory = ".notes"
const fileName = "confg.json"

func GetDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, directory), nil
}

func GetFileAbsolutePath() (string, error) {
	dir, err := GetDirectory()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, fileName), nil
}
