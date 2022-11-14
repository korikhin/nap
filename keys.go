package main

import "github.com/charmbracelet/bubbles/key"

// KeyMap is the mappings of actions to key bindings.
type KeyMap struct {
	Quit          key.Binding
	Search        key.Binding
	ToggleHelp    key.Binding
	NewSnippet    key.Binding
	RenameSnippet key.Binding
	DeleteSnippet key.Binding
	EditSnippet   key.Binding
	CopySnippet   key.Binding
	Confirm       key.Binding
	Cancel        key.Binding
	NextPane      key.Binding
	PreviousPane  key.Binding
}

// DefaultKeyMap is the default key map for the application.
var DefaultKeyMap = KeyMap{
	Quit:          key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "exit")),
	Search:        key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "search")),
	ToggleHelp:    key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	NewSnippet:    key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new")),
	DeleteSnippet: key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "delete")),
	EditSnippet:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit")),
	CopySnippet:   key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "copy")),
	Confirm:       key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "confirm")),
	Cancel:        key.NewBinding(key.WithKeys("N"), key.WithHelp("N", "cancel")),
	NextPane:      key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "navigate")),
	PreviousPane:  key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "navigate")),
}

// ShortHelp returns a quick help menu.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.NextPane,
		k.Search,
		k.EditSnippet,
		k.DeleteSnippet,
		k.CopySnippet,
		k.NewSnippet,
	}
}

// FullHelp returns all help options in a more detailed view.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NextPane, k.PreviousPane},
		{k.NextPane, k.PreviousPane},
		{k.NextPane, k.PreviousPane},
	}
}