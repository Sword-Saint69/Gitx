package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type TimelineModel struct {
	stats *git.TimelineStats
	quit  bool
}

func NewTimelineModel(s *git.TimelineStats) TimelineModel {
	return TimelineModel{stats: s}
}

func (m TimelineModel) Init() tea.Cmd { return nil }

func (m TimelineModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m TimelineModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" CONTRIBUTION TIMELINE ") + "\n\n")

	// Render last 12 weeks
	now := time.Now()
	// Start from a Sunday
	start := now.AddDate(0, 0, -int(now.Weekday()))
	start = start.AddDate(0, 0, -7*11) // 12 weeks ago

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	for i, dayName := range days {
		s.WriteString(SubtleStyle.Width(5).Render(dayName))
		
		for w := 0; w < 12; w++ {
			curr := start.AddDate(0, 0, w*7+i)
			count := m.stats.Days[curr.Format("2006-01-02")]
			
			symbol := " . "
			style := SubtleStyle
			if count > 0 && count < 3 {
				symbol = " + "
				style = SecondaryStyle
			} else if count >= 3 && count < 6 {
				symbol = " # "
				style = PrimaryStyle
			} else if count >= 6 {
				symbol = " @ "
				style = AccentStyle
			}
			
			s.WriteString(style.Render(symbol))
		}
		s.WriteString("\n")
	}

	s.WriteString("\n" + SubtleStyle.Render("Legend: . (None) + (Low) # (Med) @ (High)"))
	s.WriteString("\n" + SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
