package ui

import (
	"fmt"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/gitx/internal/git"
)

type LOCModel struct {
	stats map[string]*git.LanguageStats
	err   error
	quit  bool
}

func NewLOCModel(stats map[string]*git.LanguageStats) LOCModel {
	return LOCModel{stats: stats}
}

func (m LOCModel) Init() tea.Cmd { return nil }

func (m LOCModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m LOCModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render("ERROR: " + m.err.Error())
	}
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" LINES OF CODE ") + "\n\n")

	// Sort languages by lines
	type entry struct {
		lang  string
		stats *git.LanguageStats
	}
	var entries []entry
	totalLines := 0
	for l, st := range m.stats {
		entries = append(entries, entry{l, st})
		totalLines += st.Lines
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].stats.Lines > entries[j].stats.Lines
	})

	// Table headers
	langWidth := 15
	filesWidth := 10
	linesWidth := 10
	pctWidth := 15

	header := lipgloss.JoinHorizontal(lipgloss.Left,
		LabelStyle.Width(langWidth).Render("Language"),
		LabelStyle.Width(filesWidth).Render("Files"),
		LabelStyle.Width(linesWidth).Render("Lines"),
		LabelStyle.Width(pctWidth).Render("Percentage"),
	)
	s.WriteString(header + "\n")
	s.WriteString(SubtleStyle.Render(strings.Repeat("─", langWidth+filesWidth+linesWidth+pctWidth)) + "\n")

	for _, e := range entries {
		pct := float64(e.stats.Lines) / float64(totalLines) * 100
		progressBar := strings.Repeat("█", int(pct/10)) + strings.Repeat("░", 10-int(pct/10))

		row := lipgloss.JoinHorizontal(lipgloss.Left,
			ValueStyle.Width(langWidth).Render(e.lang),
			ValueStyle.Width(filesWidth).Render(fmt.Sprintf("%d", e.stats.Files)),
			ValueStyle.Width(linesWidth).Render(fmt.Sprintf("%d", e.stats.Lines)),
			AccentStyle.Render(fmt.Sprintf("%s %5.1f%%", progressBar, pct)),
		)
		s.WriteString(row + "\n")
	}

	s.WriteString("\n" + SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
