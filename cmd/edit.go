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
			return
		}

		notes, err := notes.List(configurations.NotesLocation)
		if err != nil {
			fmt.Println("unable to retrieve notes: " + err.Error())
			return
		}

		program := tea.NewProgram(ui.InitialModel(notes))
		if _, err := program.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
