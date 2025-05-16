package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
)

type Configuration struct {
	NotesLocation     string `json:"NotesLocation"`
	FileExtension     string `json:"FileExtension"`
	TextEditorCommand string `json:"TextEditorCommand"`
}

const notesLocationKey = "NotesLocation"
const defaultFileExtension string = "md"
const defaultTextEditorCommand string = "vi"

func Get() (*Configuration, error) {
	filepath, err := GetFileAbsolutePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, errors.New("configuration file does not exist")
	}

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	config := Configuration{
		FileExtension:     defaultFileExtension,
		TextEditorCommand: defaultTextEditorCommand,
	}

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	config.FileExtension = strings.ReplaceAll(config.FileExtension, ".", "")

	// TODO: Validate that the directory is valid
	// TODO: Make sure the directory is the absolute value

	return &config, nil
}
