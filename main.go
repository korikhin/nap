package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/maps"
)

func main() {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Println("Unable to unmarshal config", err)
	}

	var snippets []Snippet
	dir, err := os.ReadFile(config.Home + "/" + config.File)
	err = json.Unmarshal(dir, &snippets)
	if err != nil {
		fmt.Println("Unable to unmarshal snippets.json file", err)
		return
	}

	var folders = make(map[string]int)
	var items []list.Item
	for _, snippet := range snippets {
		folders[snippet.Folder]++
		items = append(items, list.Item(snippet))
	}
	snippetList := list.New(items, snippetDelegate{}, 0, 0)

	var folderItems []list.Item
	for _, folder := range maps.Keys(folders) {
		folderItems = append(folderItems, list.Item(Folder(folder)))
	}
	folderList := list.New(folderItems, folderDelegate{}, 0, 0)
	folderList.Title = "Folders"

	folderList.SetShowHelp(false)
	folderList.SetFilteringEnabled(false)
	folderList.SetShowStatusBar(false)
	folderList.DisableQuitKeybindings()

	snippetList.SetShowHelp(false)
	snippetList.SetShowFilter(true)
	snippetList.Title = "Snippets"

	snippetList.FilterInput.Prompt = "Find: "
	snippetList.FilterInput.CursorStyle = lipgloss.NewStyle().Foreground(primaryColor)
	snippetList.FilterInput.PromptStyle = lipgloss.NewStyle().Foreground(white).MarginLeft(1)
	snippetList.FilterInput.TextStyle = lipgloss.NewStyle().Foreground(white).Background(primaryColorSubdued)
	snippetList.SetStatusBarItemName("snippet", "snippets")
	snippetList.DisableQuitKeybindings()

	content := viewport.New(80, 0)

	m := &Model{
		Snippets:     snippets,
		List:         snippetList,
		Folders:      folderList,
		Code:         content,
		ContentStyle: DefaultStyles.Content.Blurred,
		ListStyle:    DefaultStyles.Snippets.Focused,
		FoldersStyle: DefaultStyles.Folders.Blurred,
		keys:         DefaultKeyMap,
		help:         help.New(),
		config:       config,
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		fmt.Println("Alas, there was an error.", err)
		return
	}
}
