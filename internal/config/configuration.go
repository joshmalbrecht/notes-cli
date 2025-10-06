package config

import (
	"encoding/json"
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

func Exists() (bool, error) {
	filepath, err := GetFileAbsolutePath()
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func Get() (*Configuration, error) {
	filepath, err := GetFileAbsolutePath()
	if err != nil {
		return nil, err
	}

	exists, err := Exists()
	if err != nil {
		return nil, err
	}

	if !exists {
		Initialize()
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
