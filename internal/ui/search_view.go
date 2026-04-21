package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type SearchModel struct {
	results []git.SearchResult
	query   string
	quit    bool
}

func NewSearchModel(query string, r []git.SearchResult) SearchModel {
	return SearchModel{query: query, results: r}
}

func (m SearchModel) Init() tea.Cmd { return nil }

func (m SearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m SearchModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" SEARCH RESULTS ") + "\n\n")
	s.WriteString(SubtleStyle.Render(fmt.Sprintf("Query: %s", m.query)) + "\n\n")

	if len(m.results) == 0 {
		s.WriteString("No matches found in history.")
	} else {
		for i, res := range m.results {
			if i > 50 { break }
			hash := AccentStyle.Render(res.Hash[:7])
			file := PrimaryStyle.Render(res.File)
			line := SubtleStyle.Render(res.Line)
			
			s.WriteString(fmt.Sprintf("%s %s:%s\n", hash, file, line))
			s.WriteString("  " + ValueStyle.Render(res.Content) + "\n\n")
		}
	}

	s.WriteString(SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
