package config

import "os"

const Directory = ".notes"
const FileName = "confg.json"

func Initialize() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filepath := homeDir + "/" + Directory

	err = os.MkdirAll(filepath, 0755)
	if err != nil {
		return err
	}

	fullFileName := filepath + "/" + FileName

	if _, err := os.Stat(fullFileName); os.IsNotExist(err) {
		file, err := os.Create(fullFileName)
		if err != nil {
			return err
		}

		_, err = file.WriteString("{\n    \"NotesLocation\":\"~/Documents\"\n}")
		if err != nil {
			return nil
		}

		println(fullFileName + " has been created")
	}

	return nil
}
