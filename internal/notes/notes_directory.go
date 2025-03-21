package notes

import (
	"os"
	paths "path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const NotesDirName = "notes-cli"

func createNewNoteFile(filepath string, title string) (string, error) {
	monthDir, err := getMonthDirectory(filepath)
	if err != nil {
		return "", nil
	}

	filename := getNoteFileName(title)
	filename = paths.Join(monthDir, filename)

	err = os.WriteFile(filename, []byte{}, 0666)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func getMonthDirectory(filepath string) (string, error) {
	monthDir := paths.Join(filepath, NotesDirName, strconv.Itoa(time.Now().Year()), strings.ToLower(time.Now().Month().String()))
	err := os.MkdirAll(monthDir, 0755)
	if err != nil {
		return "", err
	}

	return monthDir, nil
}

func getNoteFileName(title string) string {
	formattedTitle := strings.ToLower(title)
	formattedTitle = replaceSpaces(formattedTitle, "-")

	dateString := time.Now().Format("02012006")

	return dateString + "-" + formattedTitle + ".txt"
}

func replaceSpaces(s string, char string) string {
	regex := regexp.MustCompile(`\s+`)
	return regex.ReplaceAllString(s, char)
}
