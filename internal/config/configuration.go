package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const notesLocationKey = "NotesLocation"

type Configuration struct {
	NotesLocation string `json:"NotesLocation"`
}

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

	byteValue, _ := io.ReadAll(jsonFile)
	var config Configuration

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	// TODO: Validate that the directory is valid
	// TODO: Make sure the directory is the absolute value

	return &config, nil
}
