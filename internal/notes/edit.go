package notes

import (
	"os"
	"os/exec"

	"github.com/joshmalbrecht/note/internal/config"
)

// Edit will open an editor for the file with the provided name.
func Edit(filename string) error {
	filepath, err := getNoteAbsoluteFilePath(filename)
	if err != nil {
		return err
	}

	configurations, err := config.Get()
	if err != nil {
		return err
	}

	cmd := exec.Command(configurations.TextEditorCommand, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
