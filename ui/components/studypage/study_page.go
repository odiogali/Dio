package studypage

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	common "github.com/odiogali/Dio/ui/common"
)

type StudyPageModel struct {
	content   string
	ready     bool
	StudyPage viewport.Model
	KeyMap    common.KeyMap
}

func New(content string) StudyPageModel {
	return StudyPageModel{
		content: content,
		KeyMap:  common.DefaultKeyMap(),
	}
}

func (m StudyPageModel) Height() int {
	return lipgloss.Height(m.View())
}

func (m StudyPageModel) Init() tea.Cmd {
	return nil
}

func (m StudyPageModel) Update(msg tea.Msg) (StudyPageModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg: // For handling keyboard input
		switch {
		case key.Matches(msg, m.KeyMap.GotoTop):
			m.StudyPage.GotoTop()
		case key.Matches(msg, m.KeyMap.GotoBottom):
			m.StudyPage.GotoBottom()
		case key.Matches(msg, m.KeyMap.PageUp):
			m.StudyPage.PageUp()
		case key.Matches(msg, m.KeyMap.PageDown):
			m.StudyPage.PageDown()
		}
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		if !m.ready {
			m.StudyPage.YPosition = 0
			m.StudyPage.SetContent(m.content)
			m.ready = true
		}
	}

	m.StudyPage, cmd = m.StudyPage.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m StudyPageModel) View() string {
	if !m.ready {
		return "\n Initializing..."
	}

	return m.StudyPage.View()
}

func (m *StudyPageModel) SetSize(width, height int) {
	m.StudyPage.Width = width
	m.StudyPage.Height = height
}
