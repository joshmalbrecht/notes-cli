package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Choices       []string
	Cursor        int
	SelectedIndex int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			m.SelectedIndex = int(m.Cursor)
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *Model) View() string {
	view := "Select a note:\n\n"

	for i, choice := range m.Choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.Cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		view += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	view += "\nPress q to quit.\n"

	// Send the UI for rendering
	return view
}

func InitialModel(files []string) Model {
	return Model{
		Choices:       files,
		SelectedIndex: -1,
	}
}
