package notes

import (
	"os"
	"os/exec"
)

// Edit will open an editor for the file with the provided name.
func Edit(filename string) error {
	filepath, err := getNoteAbsoluteFilePath(filename)
	if err != nil {
		return err
	}

	cmd := exec.Command("vi", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
