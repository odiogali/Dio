package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/odiogali/Dio/ui/components/footer"
	"github.com/odiogali/Dio/ui/components/mappage"
	"github.com/odiogali/Dio/ui/components/navbar"
	"github.com/odiogali/Dio/ui/components/quizpage"
	"github.com/odiogali/Dio/ui/components/studypage"
)

type Model struct {
	navbar    navbar.NavbarModel
	StudyPage studypage.StudyPageModel
	QuizPage  quizpage.QuizPageModel
	MapPage   mappage.MapPageModel
	footer    footer.FooterModel
	width     int
	height    int
}

func NewModel(content string) Model {
	return Model{
		navbar:    navbar.New(),
		StudyPage: studypage.New(content),
		QuizPage:  quizpage.New(),
		footer:    footer.New(),
	}
}

func (m Model) recalculateLayout() Model {
	navbarHeight := m.navbar.Height()
	footerHeight := m.footer.Height()

	availableHeight := m.height - navbarHeight - footerHeight
	availableHeight = max(0, availableHeight)

	m.StudyPage.SetSize(m.width, availableHeight)
	m.QuizPage.Input.SetSize(int(float64(m.width)*0.35)-4, availableHeight-6)
	m.QuizPage.Output.SetSize(int(float64(m.width)*0.65), availableHeight)
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.navbar.Init(),
		m.StudyPage.Init(),
		m.footer.Init(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Update navbar and footer *first*
		m.navbar, cmd = m.navbar.Update(msg)
		cmds = append(cmds, cmd)

		m.footer, cmd = m.footer.Update(msg)
		cmds = append(cmds, cmd)

		// Now measure their rendered height (after update)
		header := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(m.navbar.View())
		footer := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(m.footer.View())
		navbarHeight := lipgloss.Height(header)
		footerHeight := lipgloss.Height(footer)

		studyPageComponentHeight := m.height - navbarHeight - footerHeight
		quizComponentHeight := m.height - navbarHeight - footerHeight

		// Study page intial sizing
		m.StudyPage.SetSize(m.width, studyPageComponentHeight)
		m.StudyPage, cmd = m.StudyPage.Update(msg)
		cmds = append(cmds, cmd)

		// Quiz page components intial sizing
		m.QuizPage.Input.SetSize(int(float64(m.width)*0.35)-4, quizComponentHeight-6)
		m.QuizPage.Output.SetSize(int(float64(m.width)*0.65), quizComponentHeight)
		m.QuizPage, cmd = m.QuizPage.Update(msg)
		cmds = append(cmds, cmd)

	case footer.HelpToggledMsg:
		return m.recalculateLayout(), tea.Batch(cmds...)

	default:
		// For other messages, just pass them through
		m.navbar, cmd = m.navbar.Update(msg)
		cmds = append(cmds, cmd)

		m.footer, cmd = m.footer.Update(msg)
		cmds = append(cmds, cmd)

		if m.navbar.ActiveTab == 0 {
			m.StudyPage, cmd = m.StudyPage.Update(msg)
		} else if m.navbar.ActiveTab == 1 {
			m.QuizPage, cmd = m.QuizPage.Update(msg)
		} else {
			m.MapPage, cmd = m.MapPage.Update(msg)
		}

		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	header := lipgloss.NewStyle().Align(lipgloss.Left).Width(m.width).Render(m.navbar.View())
	footer := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(m.footer.View())
	studyPage := lipgloss.NewStyle().
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)).
		Align(lipgloss.Center, lipgloss.Center).
		Width(m.width).
		Render(m.StudyPage.View())

	if m.navbar.ActiveTab == 0 {
		return lipgloss.JoinVertical(lipgloss.Top, header, studyPage, footer)
	}

	if m.navbar.ActiveTab == 1 {
		return lipgloss.JoinVertical(lipgloss.Top, header, m.QuizPage.View(), footer)
	}

	return lipgloss.JoinVertical(lipgloss.Top, header, m.MapPage.View(), footer)
}
