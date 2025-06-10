package quizpage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type OllamaResponseMsg string

type QuizPageModel struct {
	Input  InputModel
	Output OutputModel
}

func New() QuizPageModel {
	return QuizPageModel{
		Input:  InitialInputModel(),
		Output: InitialOutputModel(),
	}
}

func (m QuizPageModel) Init() tea.Cmd {
	return tea.Batch(
		m.Input.Init(),
		m.Output.Init(),
	)
}

func (m QuizPageModel) Update(msg tea.Msg) (QuizPageModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.Input, cmd = m.Input.Update(msg)
	cmds = append(cmds, cmd)

	m.Output, cmd = m.Output.Update(msg)
	cmds = append(cmds, cmd)

	return m, cmd
}

func (m QuizPageModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.Output.View(), m.Input.View())
}
