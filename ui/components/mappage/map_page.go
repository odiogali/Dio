package mappage

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MapPageModel struct {
}

func (m MapPageModel) Init() tea.Cmd {
	return nil
}

func (m MapPageModel) Update(msg tea.Msg) (MapPageModel, tea.Cmd) {
	return MapPageModel{}, nil
}

func (m MapPageModel) View() string {
	return ""
}
