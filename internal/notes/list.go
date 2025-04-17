package notes

import (
	"os"
	paths "path/filepath"
	"time"
)

func List(notesFilePath string) ([]string, error) {
	notesDirPath := paths.Join(notesFilePath, notesDirName)

	notesToModified := make(map[string]time.Time)
	walkDir(notesDirPath, &notesToModified)

	// TODO: Sort

	keys := make([]string, 0, len(notesToModified))
	for key := range notesToModified {
		keys = append(keys, key)
	}

	return keys, nil
}

// TODO: Rename
func walkDir(path string, notes *map[string]time.Time) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := paths.Join(path, entry.Name())

		if entry.IsDir() {
			if err := walkDir(fullPath, notes); err != nil {
				return err
			}
		} else {
			fileInfo, err := entry.Info()
			if err != nil {
				return err
			}

			(*notes)[entry.Name()] = fileInfo.ModTime()
		}
	}

	return nil
}
