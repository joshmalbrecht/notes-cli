package notes

import (
	"errors"
	"os"
	"os/exec"
)

func Create(filepath string, title string) (string, error) {
	if len(filepath) == 0 {
		return "", errors.New("filepath is empty")
	}

	if len(title) == 0 {
		return "", errors.New("title is empty")
	}

	filename, err := createNewNoteFile(filepath, title)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("vi", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		os.Remove(filename)
		return "", err
	}

	return filename, err
}
