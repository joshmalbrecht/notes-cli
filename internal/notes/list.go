package notes

import (
	"os"
	paths "path/filepath"
	"sort"
	"time"
)

// List returns the files in the provided path in descending order based on the file's modified time
func List(notesFilePath string) ([]string, error) {
	notesDirPath := paths.Join(notesFilePath, notesDirName)

	noteToModified := make(map[string]time.Time)
	walkDir(notesDirPath, &noteToModified)

	type keyValue struct {
		Key   string
		Value time.Time
	}

	// sort in descending order based on the modifed times

	var keyValues []keyValue
	for key, value := range noteToModified {
		keyValues = append(keyValues, keyValue{key, value})
	}

	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].Value.After(keyValues[j].Value)
	})

	result := make([]string, len(keyValues))
	for i, kv := range keyValues {
		result[i] = kv.Key
	}

	return result, nil
}

// walkDir recursively traverses the directories and populates the provided map with the file name and the
// file's modified time
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
