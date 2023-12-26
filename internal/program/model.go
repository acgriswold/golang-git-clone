package program

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GitModel struct {
	nameInput string
	listInput string
	event     string
}

func initialModel() (*GitModel, error) {
	return &GitModel{}, nil
}

func (m GitModel) Init() tea.Cmd {
	return nil
}

func (m GitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.Type == tea.KeyEsc {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m GitModel) View() string {
	s := "Do something. When you're done press q to quit.\n"

	return s
}

var _ tea.Model = &GitModel{}
