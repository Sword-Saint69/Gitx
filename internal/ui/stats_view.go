package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/gitx/internal/git"
)

type StatsModel struct {
	stats    *git.RepoStats
	err      error
	quitting bool
}

func NewStatsModel(stats *git.RepoStats) StatsModel {
	return StatsModel{
		stats: stats,
	}
}

func (m StatsModel) Init() tea.Cmd {
	return nil
}

func (m StatsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m StatsModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render("ERROR: " + m.err.Error())
	}
	if m.quitting { return "" }

	var s strings.Builder

	// Top Section
	s.WriteString(HeaderStyle.Render(" REPOSITORY INSIGHTS ") + "\n\n")

	// High Level Metrics
	metrics := lipgloss.JoinHorizontal(lipgloss.Top,
		PanelStyle.Render(fmt.Sprintf("%s\n%s", PrimaryStyle.Render("COMMITS"), ValueStyle.Render(fmt.Sprintf("%d", m.stats.Commits)))),
		PanelStyle.Render(fmt.Sprintf("%s\n%s", PrimaryStyle.Render("AUTHORS"), ValueStyle.Render(fmt.Sprintf("%d", m.stats.Contributors)))),
		PanelStyle.Render(fmt.Sprintf("%s\n%s", PrimaryStyle.Render("BRANCHES"), ValueStyle.Render(fmt.Sprintf("%d", m.stats.Branches)))),
	)
	s.WriteString(metrics + "\n\n")

	// Activity Summary Card
	activityInfo := fmt.Sprintf("%s\n%s",
		InfoField("Active Since:", m.stats.Age),
		InfoField("Last Commit: ", m.stats.LastCommit.Format("2006-01-02 15:04")),
	)
	
	s.WriteString(Card("Activity Summary", activityInfo) + "\n\n")
	s.WriteString(SubtleStyle.Render("Press 'q' to exit"))

	return s.String()
}
