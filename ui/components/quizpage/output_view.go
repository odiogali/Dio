package quizpage

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/glamour"
	common "github.com/odiogali/Dio/ui/common"
)

type OutputModel struct {
	content        string
	ready          bool
	OutputViewport viewport.Model
	loading        common.SpinnerModel
	send           func(tea.Msg)
}

func InitialOutputModel() OutputModel {
	msg := Message{
		Role:    "user",
		Content: "Why is the sky blue",
	}

	_ = Request{
		Model:    "llama3.2",
		Stream:   true,
		Messages: []Message{msg},
	}

	// ch, errc := StreamChunks(DefaultOllamaURL, req)
	// go func() {
	// 	for {
	// 		select {
	// 		case chunk, ok := <-ch:
	// 			if !ok {
	// 				return
	// 			}
	// 			program.send(OllamaResponseMsg(chunk))
	// 		case err := <-errc:
	// 			if err != nil {
	// 				program.send(OllamaResponseMsg("[error: " + err.Error() + "]"))
	// 				return
	// 			}
	// 		}
	// 	}
	// }()

	return OutputModel{
		content: "",
	}
}

func (m OutputModel) Init() tea.Cmd {
	return nil
}

func (m OutputModel) Update(msg tea.Msg) (OutputModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.OutputViewport.YPosition = 0
			m.OutputViewport.SetContent(m.content)

			m.ready = true
		}

	case OllamaResponseMsg:
		// Set viewport content to new AI response
		m.content += string(msg)
		m.OutputViewport.SetContent(m.content)
	}

	return m, cmd
}

func (m OutputModel) View() string {
	if !m.ready {
		return "Initializing..."
	}

	return m.OutputViewport.View()
}

func (m *OutputModel) SetSize(width, height int) {
	m.OutputViewport.Width = width
	m.OutputViewport.Height = height
}
