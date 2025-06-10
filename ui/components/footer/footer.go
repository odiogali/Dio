package footer

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	common "github.com/odiogali/Dio/ui/common"
)

type FooterModel struct {
	keys       common.KeyMap
	help       help.Model
	inputStyle lipgloss.Style
	width      int
}

func New() FooterModel {
	return FooterModel{
		keys:       common.DefaultKeyMap(),
		help:       help.New(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
	}
}

type HelpToggledMsg struct{}

func (m FooterModel) Init() tea.Cmd {
	return nil
}

func (m FooterModel) Update(msg tea.Msg) (FooterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
		m.width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, func() tea.Msg { return HelpToggledMsg{} }
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m FooterModel) Height() int {
	return lipgloss.Height(m.View())
}

func (m FooterModel) View() string {
	helpView := m.help.View(m.keys)

	divider := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#757b94")).
		SetString(strings.Repeat("─", m.width)).Render()

	return divider + "\n" + helpView
}
