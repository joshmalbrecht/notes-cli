package notes

import (
	"fmt"
	"os"
	"path/filepath"
	paths "path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joshmalbrecht/note/internal/config"
	"github.com/joshmalbrecht/note/internal/util"
)

const notesDirName string = "notes-cli"
const dateFormat string = "02012006"

func createNewNoteFile(filepath string, title string, fileExtension string) (string, error) {
	monthDir, err := getMonthDirectory(filepath)
	if err != nil {
		return "", nil
	}

	filename := getNoteFileName(title, fileExtension)
	filename = paths.Join(monthDir, filename)

	err = os.WriteFile(filename, []byte{}, 0666)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func getMonthDirectory(filepath string) (string, error) {
	monthDir := paths.Join(filepath, notesDirName, strconv.Itoa(time.Now().Year()), strings.ToLower(time.Now().Month().String()))
	err := os.MkdirAll(monthDir, 0755)
	if err != nil {
		return "", err
	}

	return monthDir, nil
}

func getNoteFileName(title string, fileExtension string) string {
	formattedTitle := strings.ToLower(title)
	formattedTitle = util.ReplaceSpaces(formattedTitle, "-")

	dateString := time.Now().Format(dateFormat)

	return dateString + "-" + formattedTitle + "." + fileExtension
}

// getNoteAbsoluteFilePath searches for the absolute file path for the file with the povided name. If the file
// cannot be found, an error will be returned.
func getNoteAbsoluteFilePath(filename string) (string, error) {
	configurations, err := config.Get()
	if err != nil {
		return "", err
	}

	rootNotesDir := paths.Join(configurations.NotesLocation, notesDirName)
	err = os.MkdirAll(rootNotesDir, 0755)
	if err != nil {
		return "", err
	}

	// TODO: Now find the file in the root notes directory
	filePath, err := findFile(rootNotesDir, filename)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// findFile searches the root directory and all of the children directories for the provided filename. The
// absolute filepath for the file will be returned if it is found or an error will be returned if the file
// cannot be found.
func findFile(rootDir string, targetFilename string) (string, error) {
	var foundPath string
	err := filepath.WalkDir(rootDir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !entry.IsDir() && entry.Name() == targetFilename {
			foundPath = path
			// Returning an error to break early once we found the file
			return fmt.Errorf("found")
		}

		return nil
	})

	if foundPath != "" {
		return foundPath, nil
	}

	if err != nil && err.Error() != "found" {
		return foundPath, err
	}

	return "", fmt.Errorf("file %q not found under %q", targetFilename, rootDir)
}
