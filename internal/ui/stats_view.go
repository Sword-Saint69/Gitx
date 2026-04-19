package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
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
		return ErrorColor.Render(fmt.Sprintf("Error: %v", m.err))
	}

	if m.quitting {
		return ""
	}

	// Build the view
	var s strings.Builder

	// Header
	s.WriteString(HeaderStyle.Render(" 🔀 GITX - Repository Insights ") + "\n\n")

	// Stats Grid
	commitsBox := StatBoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			LabelStyle.Render("Commits"),
			ValueStyle.Bold(true).Render(fmt.Sprintf("%d", m.stats.Commits)),
		),
	)

	authorsBox := StatBoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			LabelStyle.Render("Contributors"),
			ValueStyle.Bold(true).Render(fmt.Sprintf("%d", m.stats.Contributors)),
		),
	)

	branchesBox := StatBoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			LabelStyle.Render("Branches"),
			ValueStyle.Bold(true).Render(fmt.Sprintf("%d", m.stats.Branches)),
		),
	)

	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, commitsBox, authorsBox, branchesBox) + "\n\n")

	// Info section
	info := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("%s %s", LabelStyle.Render("Active Since:"), ValueStyle.Render(m.stats.Age)),
		fmt.Sprintf("%s %s", LabelStyle.Render("Last Commit: "), ValueStyle.Render(m.stats.LastCommit.Format("2006-01-02 15:04"))),
	)

	s.WriteString(BorderStyle.Render(info) + "\n\n")

	s.WriteString(SubtleColor.Render("Press 'q' to exit"))

	return s.String()
}
