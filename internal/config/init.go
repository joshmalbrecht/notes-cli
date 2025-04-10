package config

import (
	"fmt"
	"os"
)

func Initialize() error {
	filepath, err := GetDirectory()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath, 0755)
	if err != nil {
		return err
	}

	fullFileName := filepath + "/" + fileName

	if _, err := os.Stat(fullFileName); os.IsNotExist(err) {
		file, err := os.Create(fullFileName)
		if err != nil {
			return err
		}

		_, err = file.WriteString("{\n    \"" + notesLocationKey + "\":\"~/Documents\"\n}")
		if err != nil {
			return nil
		}

		fmt.Println(fullFileName + " has been created")
	}

	return nil
}
