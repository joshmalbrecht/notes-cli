package config

import (
	"fmt"
	"os"
	"os/user"
)

func Initialize() error {
	fmt.Println("initializing the configuration file ...")

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

		user, err := user.Current()
		if err != nil {
			return err
		}

		_, err = file.WriteString("{\n    \"" + notesLocationKey + "\":\"" + user.HomeDir + "\"\n}")
		if err != nil {
			return nil
		}

		fmt.Println(fullFileName + " has been created")
	}

	return nil
}
