package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type HotspotsModel struct {
	hotspots []git.Hotspot
	quit     bool
}

func NewHotspotsModel(h []git.Hotspot) HotspotsModel {
	return HotspotsModel{hotspots: h}
}

func (m HotspotsModel) Init() tea.Cmd { return nil }

func (m HotspotsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m HotspotsModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" HOTSPOT ANALYSIS ") + "\n\n")
	s.WriteString(SubtleStyle.Render("High Churn + High Complexity = Risk") + "\n\n")

	for _, h := range m.hotspots {
		riskLevel := "LOW"
		style := SecondaryStyle
		if h.Score > 5000 {
			riskLevel = "CRITICAL"
			style = ErrorStyle
		} else if h.Score > 2000 {
			riskLevel = "HIGH"
			style = WarningStyle
		}

		s.WriteString(fmt.Sprintf("%-25s %s\n", PrimaryStyle.Render(h.File), style.Render("["+riskLevel+"]")))
		s.WriteString(fmt.Sprintf("  %s %d | %s %d\n\n", SubtleStyle.Render("Churn:"), h.Churn, SubtleStyle.Render("LOC:"), h.LOC))
	}

	s.WriteString(SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
