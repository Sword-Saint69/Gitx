package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/gitx/internal/git"
)

type ContributorsModel struct {
	contributors []git.Contributor
	quit         bool
}

func NewContributorsModel(c []git.Contributor) ContributorsModel {
	return ContributorsModel{contributors: c}
}

func (m ContributorsModel) Init() tea.Cmd { return nil }

func (m ContributorsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m ContributorsModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" 👥 Top Contributors ") + "\n\n")

	nameWidth := 25
	countWidth := 10
	barWidth := 20

	header := lipgloss.JoinHorizontal(lipgloss.Left,
		LabelStyle.Width(nameWidth).Render("Author"),
		LabelStyle.Width(countWidth).Render("Commits"),
		LabelStyle.Width(barWidth).Render("Activity"),
	)
	s.WriteString(header + "\n")
	s.WriteString(SubtleStyle.Render(strings.Repeat("─", nameWidth+countWidth+barWidth)) + "\n")

	if len(m.contributors) == 0 {
		return s.String() + "No contributors found."
	}

	maxCommits := m.contributors[0].Commits

	for _, c := range m.contributors {
		barLen := 0
		if maxCommits > 0 {
			barLen = (c.Commits * barWidth) / maxCommits
		}
		bar := SecondaryStyle.Render(strings.Repeat("■", barLen))

		row := lipgloss.JoinHorizontal(lipgloss.Left,
			ValueStyle.Width(nameWidth).Render(c.Name),
			ValueStyle.Width(countWidth).Render(fmt.Sprintf("%d", c.Commits)),
			bar,
		)
		s.WriteString(row + "\n")
	}

	s.WriteString("\n" + SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
