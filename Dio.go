package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/odiogali/Dio/ui"
)

func main() {
	in, err := os.ReadFile("Haskell.md")
	if err != nil {
		fmt.Println("could not load file: ", err)
		os.Exit(1)
	}

	out, err := glamour.Render(string(in), "dracula")
	if err != nil {
		fmt.Println("unable to render the markdown content: ", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		ui.NewModel(out),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program: ", err)
		os.Exit(1)
	}
}
