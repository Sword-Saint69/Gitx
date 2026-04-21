package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type AnalysisModel struct {
	mode     string
	churn    []git.ChurnEntry
	bus      []git.BusFactorEntry
	quit     bool
}

func NewChurnModel(entries []git.ChurnEntry) AnalysisModel {
	return AnalysisModel{mode: "CHURN", churn: entries}
}

func NewBusFactorModel(entries []git.BusFactorEntry) AnalysisModel {
	return AnalysisModel{mode: "BUS FACTOR", bus: entries}
}

func (m AnalysisModel) Init() tea.Cmd { return nil }

func (m AnalysisModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m AnalysisModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" "+m.mode+" ") + "\n\n")

	if m.mode == "CHURN" {
		s.WriteString(SubtleStyle.Render("Files with the most changes (bug magnets)") + "\n\n")
		for _, e := range m.churn {
			bar := PrimaryStyle.Render(strings.Repeat("█", e.Count/5+1))
			s.WriteString(fmt.Sprintf("%-30s %s %d\n", e.File, bar, e.Count))
		}
	} else {
		s.WriteString(SubtleStyle.Render("Files with low contributor count") + "\n\n")
		for i, e := range m.bus {
			if i > 20 { break } // Top 20 risky files
			style := SecondaryStyle
			if e.Authors == 1 {
				style = ErrorStyle
			}
			s.WriteString(fmt.Sprintf("%-30s %s %d authors\n", e.File, style.Render(fmt.Sprintf("%d", e.Authors)), e.Authors))
		}
	}

	s.WriteString("\n" + SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
