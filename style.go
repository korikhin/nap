package main

import "github.com/charmbracelet/lipgloss"

// SnippetsStyle is the style struct to handle the focusing and blurring of the
// snippets pane in the application.
type SnippetsStyle struct {
	Focused SnippetsBaseStyle
	Blurred SnippetsBaseStyle
}

// FoldersStyle is the style struct to handle the focusing and blurring of the
// folders pane in the application.
type FoldersStyle struct {
	Focused FoldersBaseStyle
	Blurred FoldersBaseStyle
}

// ContentStyle is the style struct to handle the focusing and blurring of the
// content pane in the application.
type ContentStyle struct {
	Focused ContentBaseStyle
	Blurred ContentBaseStyle
}

// SnippetsBaseStyle holds the neccessary styling for the snippets pane of
// the application.
type SnippetsBaseStyle struct {
	Base               lipgloss.Style
	Title              lipgloss.Style
	TitleBar           lipgloss.Style
	StatusBar          lipgloss.Style
	SelectedSubtitle   lipgloss.Style
	UnselectedSubtitle lipgloss.Style
	SelectedTitle      lipgloss.Style
	UnselectedTitle    lipgloss.Style
	CopiedTitleBar     lipgloss.Style
	CopiedTitle        lipgloss.Style
	CopiedSubtitle     lipgloss.Style
	DeletedTitleBar    lipgloss.Style
	DeletedTitle       lipgloss.Style
	DeletedSubtitle    lipgloss.Style
}

// FoldersBaseStyle holds the neccessary styling for the folders pane of
// the application.
type FoldersBaseStyle struct {
	Base       lipgloss.Style
	Title      lipgloss.Style
	TitleBar   lipgloss.Style
	Selected   lipgloss.Style
	Unselected lipgloss.Style
}

// ContentBaseStyle holds the neccessary styling for the content pane of the
// application.
type ContentBaseStyle struct {
	Base         lipgloss.Style
	Title        lipgloss.Style
	Separator    lipgloss.Style
	LineNumber   lipgloss.Style
	EmptyHint    lipgloss.Style
	EmptyHintKey lipgloss.Style
}

// Styles is the struct of all styles for the application.
type Styles struct {
	Snippets SnippetsStyle
	Folders  FoldersStyle
	Content  ContentStyle
}

var marginStyle = lipgloss.NewStyle().Margin(1, 0, 0, 1)

// DefaultStyles is the default implementation of the styles struct for all
// styling in the application.
func DefaultStyles(config Config) Styles {
	background := lipgloss.Color(config.BackgroundColor)
	black := lipgloss.Color(config.BlackColor)
	gray := lipgloss.Color(config.GrayColor)
	brightGray := lipgloss.Color(config.BrightGrayColor)
	white := lipgloss.Color(config.WhiteColor)
	red := lipgloss.Color(config.RedColor)
	brightRed := lipgloss.Color(config.BrightRedColor)
	green := lipgloss.Color(config.GreenColor)
	brightGreen := lipgloss.Color(config.BrightGreenColor)
	blue := lipgloss.Color(config.PrimaryColorSubdued)
	brightBlue := lipgloss.Color(config.PrimaryColor)

	// Extra styles
	var status lipgloss.Color
	if s := config.StatusColor; s != "" {
		status = lipgloss.Color(s)
	} else {
		status = gray
	}

	return Styles{
		Snippets: SnippetsStyle{
			Focused: SnippetsBaseStyle{
				Base:               lipgloss.NewStyle().Width(35),
				Title:              lipgloss.NewStyle().Foreground(white),
				TitleBar:           lipgloss.NewStyle().Background(blue).Foreground(white).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				StatusBar:          lipgloss.NewStyle().Foreground(status).MaxWidth(35-2).Margin(1, 2),
				SelectedTitle:      lipgloss.NewStyle().Foreground(brightBlue),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(brightGray),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(blue),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(gray),
				CopiedTitle:        lipgloss.NewStyle().Foreground(brightGreen),
				CopiedTitleBar:     lipgloss.NewStyle().Background(green).Foreground(white).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				CopiedSubtitle:     lipgloss.NewStyle().Foreground(green),
				DeletedTitle:       lipgloss.NewStyle().Foreground(brightRed),
				DeletedTitleBar:    lipgloss.NewStyle().Background(red).Foreground(white).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				DeletedSubtitle:    lipgloss.NewStyle().Foreground(red),
			},
			Blurred: SnippetsBaseStyle{
				Base:               lipgloss.NewStyle().Width(35),
				Title:              lipgloss.NewStyle().Foreground(brightGray),
				TitleBar:           lipgloss.NewStyle().Background(background).Foreground(brightGray).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				StatusBar:          lipgloss.NewStyle().Foreground(status).MaxWidth(35-2).Margin(1, 2),
				SelectedTitle:      lipgloss.NewStyle().Foreground(brightBlue),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(gray),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(blue),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(black),
				CopiedTitle:        lipgloss.NewStyle().Foreground(brightGreen),
				CopiedTitleBar:     lipgloss.NewStyle().Background(green).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				CopiedSubtitle:     lipgloss.NewStyle().Foreground(green),
				DeletedTitle:       lipgloss.NewStyle().Foreground(brightRed),
				DeletedTitleBar:    lipgloss.NewStyle().Background(red).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				DeletedSubtitle:    lipgloss.NewStyle().Foreground(red),
			},
		},
		Folders: FoldersStyle{
			Focused: FoldersBaseStyle{
				Base:       lipgloss.NewStyle().Width(22),
				Title:      lipgloss.NewStyle().Foreground(white).Padding(0, 1),
				TitleBar:   lipgloss.NewStyle().Background(blue).Width(22-2).Margin(0, 1, 1, 1),
				Selected:   lipgloss.NewStyle().Foreground(brightBlue),
				Unselected: lipgloss.NewStyle().Foreground(brightGray),
			},
			Blurred: FoldersBaseStyle{
				Base:       lipgloss.NewStyle().Width(22),
				Title:      lipgloss.NewStyle().Foreground(brightGray).Padding(0, 1),
				TitleBar:   lipgloss.NewStyle().Background(background).Width(22-2).Margin(0, 1, 1, 1),
				Selected:   lipgloss.NewStyle().Foreground(brightBlue),
				Unselected: lipgloss.NewStyle().Foreground(gray),
			},
		},
		Content: ContentStyle{
			Focused: ContentBaseStyle{
				Base:         lipgloss.NewStyle().Margin(0, 1),
				Title:        lipgloss.NewStyle().Background(blue).Foreground(white).Margin(0, 0, 1, 1).Padding(0, 1),
				Separator:    lipgloss.NewStyle().Foreground(white).Margin(0, 0, 1, 1),
				LineNumber:   lipgloss.NewStyle().Foreground(gray),
				EmptyHint:    lipgloss.NewStyle().Foreground(brightGray),
				EmptyHintKey: lipgloss.NewStyle().Foreground(brightBlue),
			},
			Blurred: ContentBaseStyle{
				Base:         lipgloss.NewStyle().Margin(0, 1),
				Title:        lipgloss.NewStyle().Background(background).Foreground(brightGray).Margin(0, 0, 1, 1).Padding(0, 1),
				Separator:    lipgloss.NewStyle().Foreground(brightGray).Margin(0, 0, 1, 1),
				LineNumber:   lipgloss.NewStyle().Foreground(black),
				EmptyHint:    lipgloss.NewStyle().Foreground(brightGray),
				EmptyHintKey: lipgloss.NewStyle().Foreground(brightBlue),
			},
		},
	}
}
