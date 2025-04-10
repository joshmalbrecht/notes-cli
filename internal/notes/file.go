package notes

import (
	"os"
	paths "path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joshmalbrecht/note/internal/util"
)

const notesDirName string = "notes-cli"
const dateFormat string = "02012006"

func createNewNoteFile(filepath string, title string, fileExtension string) (string, error) {
	monthDir, err := getMonthDirectory(filepath)
	if err != nil {
		return "", nil
	}

	filename := getNoteFileName(title, fileExtension)
	filename = paths.Join(monthDir, filename)

	err = os.WriteFile(filename, []byte{}, 0666)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func getMonthDirectory(filepath string) (string, error) {
	monthDir := paths.Join(filepath, notesDirName, strconv.Itoa(time.Now().Year()), strings.ToLower(time.Now().Month().String()))
	err := os.MkdirAll(monthDir, 0755)
	if err != nil {
		return "", err
	}

	return monthDir, nil
}

func getNoteFileName(title string, fileExtension string) string {
	formattedTitle := strings.ToLower(title)
	formattedTitle = util.ReplaceSpaces(formattedTitle, "-")

	dateString := time.Now().Format(dateFormat)

	return dateString + "-" + formattedTitle + "." + fileExtension
}
