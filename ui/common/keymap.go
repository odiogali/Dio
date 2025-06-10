package common

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	GotoTop    key.Binding
	GotoBottom key.Binding
	PageUp     key.Binding
	PageDown   key.Binding
	Help       key.Binding
	Quit       key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "scroll up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "scroll down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←", "move to left tab"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→", "move to right tab"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("g"),
			key.WithHelp("g", "go to top"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("G"),
			key.WithHelp("G", "go to bottom"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("ctrl+f"),
			key.WithHelp("ctrl+f", "page down"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	// Keymaps are arranged column by column (don't exceed four items per col)
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.GotoTop, k.GotoBottom, k.PageUp, k.PageDown},
		{k.Help, k.Quit},
	}
}

// var Keys = KeyMap{
// 	Up: key.NewBinding(
// 		key.WithKeys("up", "k"),
// 		key.WithHelp("↑/k", "move up"),
// 	),
// 	Down: key.NewBinding(
// 		key.WithKeys("down", "j"),
// 		key.WithHelp("↓/j", "move down"),
// 	),
// 	Left: key.NewBinding(
// 		key.WithKeys("left", "h"),
// 		key.WithHelp("←/h", "move left"),
// 	),
// 	Right: key.NewBinding(
// 		key.WithKeys("right", "l"),
// 		key.WithHelp("→/l", "move right"),
// 	),
// 	Help: key.NewBinding(
// 		key.WithKeys("?"),
// 		key.WithHelp("?", "toggle help"),
// 	),
// 	Quit: key.NewBinding(
// 		key.WithKeys("q", "ctrl+c"),
// 		key.WithHelp("q", "quit"),
// 	),
// 	GotoTop: key.NewBinding(
// 		key.WithKeys("g", "g"),
// 		key.WithHelp("gg", "go to top"),
// 	),
// 	GotoBottom: key.NewBinding(
// 		key.WithKeys("G"),
// 		key.WithHelp("G", "go to bottom"),
// 	),
// 	PageUp: key.NewBinding(
// 		key.WithKeys("ctrl", "b"),
// 		key.WithHelp("ctrl+b", "page up"),
// 	),
// 	PageDown: key.NewBinding(
// 		key.WithKeys("ctrl", "f"),
// 		key.WithHelp("ctrl+f", "page down"),
// 	),
// }
