package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joshmalbrecht/note/internal/config"
	"github.com/joshmalbrecht/note/internal/notes"
	"github.com/joshmalbrecht/note/internal/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configurations, err := config.Get()
		if err != nil {
			fmt.Println("unable to read config: " + err.Error())
			os.Exit(1)
		}

		notes, err := notes.List(configurations.NotesLocation)
		if err != nil {
			fmt.Println("unable to retrieve notes: " + err.Error())
			os.Exit(1)
		}

		model := ui.InitialModel(notes)
		program := tea.NewProgram(model)
		if _, err := program.Run(); err != nil {
			fmt.Printf("error running selection ui: %v", err)
			os.Exit(1)
		}

		var index int
		for i, val := range model.Selected {
			if val == struct{}{} {
				index = i
				break
			}
		}

		var fileName = model.Choices[index]
		fmt.Println(fileName)

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
