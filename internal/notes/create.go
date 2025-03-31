package notes

import (
	"errors"
	"os"
	"os/exec"
)

// Create creates a new note file at the provided filepath with a formatted file name based on the provided title.
// A text editor command is then run for the user to populate the note with text.
// If the note is empty after the text editor is closed, the file will be deleted.
func Create(filepath string, title string) (string, bool, error) {
	if len(filepath) == 0 {
		return "", false, errors.New("filepath is empty")
	}

	if len(title) == 0 {
		return "", false, errors.New("title is empty")
	}

	filename, err := createNewNoteFile(filepath, title)
	if err != nil {
		return "", false, err
	}

	cmd := exec.Command("vi", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		os.Remove(filename)
		return "", false, err
	}

	// Check to see if the user did not add any text to the file.
	// If the file is empty, we can remove the file.
	info, err := os.Stat(filename)
	if err != nil {
		return "", false, err
	}

	if info.Size() == 0 {
		os.Remove(filename)
		return "", false, nil
	}

	return filename, true, err
}
