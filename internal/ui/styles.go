package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Professional Noir/Oceanic Palette
	PrimaryColor   = lipgloss.Color("#00D7FF") // Cyan
	SecondaryColor = lipgloss.Color("#50FA7B") // Emerald
	AccentColor    = lipgloss.Color("#BD93f9") // Purple
	WarningColor   = lipgloss.Color("#FFB86C") // Orange
	ErrorColor     = lipgloss.Color("#FF5555") // Red
	BgColor        = lipgloss.Color("#1a1a1a") // Deep Charcoal
	FgColor        = lipgloss.Color("#F8F8F2") // Off-white
	SubtleColor    = lipgloss.Color("#44475a") // Muted Blue/Grey

	// --- Component Styles ---
	
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(PrimaryColor).
			Padding(0, 2).
			MarginBottom(1)

	PrimaryStyle = lipgloss.NewStyle().Foreground(PrimaryColor).Bold(true)
	SecondaryStyle = lipgloss.NewStyle().Foreground(SecondaryColor)
	AccentStyle = lipgloss.NewStyle().Foreground(AccentColor)
	SubtleStyle = lipgloss.NewStyle().Foreground(SubtleColor)
	ValueStyle = lipgloss.NewStyle().Foreground(FgColor).Bold(true)
	ErrorStyle = lipgloss.NewStyle().Foreground(ErrorColor).Bold(true)
	WarningStyle = lipgloss.NewStyle().Foreground(WarningColor).Bold(true)

	// Boxed Panel for content grouping
	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(SubtleColor).
			Padding(1, 2).
			MarginBottom(1)

	// --- Icons ---
	IconCheck = "✓"
	IconWarn  = "!"
	IconInfo  = "i"
	IconArrow = "»"

	// --- Compatibility Aliases ---
	TitleStyle  = HeaderStyle
	LabelStyle  = PrimaryStyle
	IconWarning = IconWarn
	IconDot     = "•"
)

func BlockHeader(title string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color("#2a2a2a")).
		Foreground(PrimaryColor).
		Bold(true).
		Padding(0, 1).
		MarginTop(1).
		Render(" " + strings.ToUpper(title) + " ")
}

func InfoField(label, value string) string {
	return fmt.Sprintf("%s %s", PrimaryStyle.Render(label), ValueStyle.Render(value))
}

func Card(title, content string) string {
	header := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(SecondaryColor).
		Padding(0, 1).
		Bold(true).
		Render(strings.ToUpper(title))
	
	return lipgloss.JoinVertical(lipgloss.Left, 
		header,
		PanelStyle.Render(content),
	)
}
