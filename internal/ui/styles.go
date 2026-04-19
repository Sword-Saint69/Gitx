package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Professional Noir/Oceanic Palette
	PrimaryColor   = lipgloss.Color("#00D7FF") // Cyan
	SecondaryColor = lipgloss.Color("#50FA7B") // Emerald
	AccentColor    = lipgloss.Color("#BD93f9") // Purple
	WarningColor   = lipgloss.Color("#FFB86C") // Orange
	ErrorColor     = lipgloss.Color("#FF5555") // Red
	BgColor        = lipgloss.Color("#282A36") // Dark Grey
	FgColor        = lipgloss.Color("#F8F8F2") // Off-white
	SubtleColor    = lipgloss.Color("#6272A4") // Muted Blue

	// Base Styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor).
			Padding(0, 1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(PrimaryColor).
			Padding(0, 2).
			MarginBottom(1).
			SetString("GITX")

	LabelStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true)

	ValueStyle = lipgloss.NewStyle().
			Foreground(FgColor)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(SubtleColor).
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

	// Icons (Text-based)
	IconArrow   = ">>>"
	IconDot     = "•"
	IconInfo    = "[i]"
	IconWarning = "[!]"
	IconCheck   = "[OK]"
)

func BlockHeader(text string) string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(SubtleColor).
		Padding(0, 1).
		Render(text)
}
