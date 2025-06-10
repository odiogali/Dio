package navbar

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(5).
			PaddingRight(5).
			Foreground(lipgloss.Color("#c1caf2")).
			Background(lipgloss.Color("#3f4767"))
	inActiveTabStyle = lipgloss.NewStyle().
				PaddingLeft(5).
				PaddingRight(5).
				Foreground(lipgloss.Color("#757b94"))
)

type NavbarModel struct {
	tabs      []string
	ActiveTab int
	width     int
}

func New() NavbarModel {
	return NavbarModel{
		// Study tab looks at the currently active review note (interactive)
		// Quiz tab is the view for the AI asking questions and evaluating correctness
		// Roadmap tab is the view for what has been reviewed, progress, and what is next
		// Settings is the view for changing what you want to review, the kinds of questions asked for the type of MD file,
		//      changes to file structure of notes, etc.
		tabs:      []string{"Study", "Quiz", "Roadmap", "Settings"},
		ActiveTab: 0,
	}
}

func (m NavbarModel) Height() int {
	return lipgloss.Height(m.View())
}

func (m NavbarModel) Init() tea.Cmd {
	return nil
}

func (m NavbarModel) Update(msg tea.Msg) (NavbarModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.ActiveTab > 0 {
				m.ActiveTab--
			}
		case "right":
			if m.ActiveTab < len(m.tabs)-1 {
				m.ActiveTab++
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m NavbarModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}
	var renderedTabs []string

	magGlassStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#757b94"))
	magGlass := magGlassStyle.Render("🔍")

	renderedTabs = append(renderedTabs, magGlass)
	for i, tab := range m.tabs {
		if i == m.ActiveTab {
			renderedTabs = append(renderedTabs, ActiveTabStyle.Render(tab))
		} else {
			renderedTabs = append(renderedTabs, inActiveTabStyle.Render(tab))
		}
	}

	separatorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#c1caf2"))
	separtor := separatorStyle.Render(" | ")

	navbar := strings.Join(renderedTabs, separtor)
	navbarStyle := lipgloss.NewStyle().
		MarginTop(1).
		MarginLeft(5).
		Width(m.width).
		Align(lipgloss.Left)

	divider := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#757b94")).
		SetString(strings.Repeat("─", m.width)).
		Render()

	return navbarStyle.Render(navbar) + "\n" + divider
}
