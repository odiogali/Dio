package quizpage

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type InputModel struct {
	textArea textarea.Model
	err      error
}

func InitialInputModel() InputModel {
	ta := textarea.New()
	ta.Placeholder = "Enter your answer here..."
	ta.Focus()
	ta.CharLimit = 0
	ta.ShowLineNumbers = false

	ta.FocusedStyle = textarea.Style{
		Prompt: lipgloss.NewStyle().Foreground(lipgloss.Color("63")),

		Placeholder: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),

		Text: lipgloss.NewStyle().Foreground(lipgloss.Color("#c1caf2")),
	}

	return InputModel{
		textArea: ta,
		err:      nil,
	}
}

func (m InputModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "?" {
			return m, nil
		}
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textArea, cmd = m.textArea.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	textInputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("#c1caf2")).
		Foreground(lipgloss.Color("#c1caf2"))

	inputContent := fmt.Sprintf(
		"\n%s\n\n%s",
		m.textArea.View(),
		"(enter to send response)",
	) + "\n"

	textInput := textInputStyle.Render(inputContent)

	return textInput
}

func (m *InputModel) SetSize(width, height int) {
	m.textArea.SetWidth(width)
	m.textArea.SetHeight(height)
}
