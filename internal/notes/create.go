package notes

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func Create(filepath string, title string) (string, error) {
	if len(filepath) == 0 {
		return "", errors.New("filepath is empty")
	}

	if len(title) == 0 {
		return "", errors.New("title is empty")
	}

	filename := createFileName(title)
	filename = filepath + filename

	err := os.WriteFile(filename, []byte{}, 0755)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("vi", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		// TODO: Delete the file
		return "", err
	}

	return filename, err
}

func createFileName(title string) string {
	formattedTitle := strings.ToLower(title)
	formattedTitle = replaceSpaces(formattedTitle, "-")

	dateString := time.Now().Format("02012006")

	return dateString + "-" + formattedTitle + ".txt"
}

func replaceSpaces(s string, char string) string {
	regex := regexp.MustCompile(`\s+`)
	return regex.ReplaceAllString(s, char)
}
