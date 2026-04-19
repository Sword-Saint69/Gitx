package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	PrimaryColor   = lipgloss.Color("#7D56F4")
	SecondaryColor = lipgloss.Color("#04B575")
	AccentColor    = lipgloss.Color("#EE6FF8")
	WarningColor   = lipgloss.Color("#FFA100")
	ErrorColor     = lipgloss.Color("#FF4C4C")
	BgColor        = lipgloss.Color("#1A1B26")
	FgColor        = lipgloss.Color("#C0CAF5")
	SubtleColor    = lipgloss.Color("#565F89")

	// Styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor).
			Padding(0, 1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(PrimaryColor).
			Padding(0, 1).
			MarginBottom(1)

	InfoStyle = lipgloss.NewStyle().
			Foreground(FgColor)

	LabelStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true)

	ValueStyle = lipgloss.NewStyle().
			Foreground(FgColor)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2)

	StatBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(SubtleColor).
			Padding(0, 1).
			MarginRight(2)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)

	SubtleStyle = lipgloss.NewStyle().
			Foreground(SubtleColor)

	PrimaryStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	SecondaryStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor)

	AccentStyle = lipgloss.NewStyle().
			Foreground(AccentColor)

	KeywordStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true)
)

func GradientText(text string) string {
	// Simple implementation for now, lipgloss doesn't have native multi-color gradients for strings easily
	// but we can simulate it or just use a nice solid bold color.
	return TitleStyle.Render(text)
}
