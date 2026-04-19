package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/scanner"
)

type SecretsModel struct {
	matches []scanner.SecretMatch
	err     error
	quit    bool
}

func NewSecretsModel(m []scanner.SecretMatch) SecretsModel {
	return SecretsModel{matches: m}
}

func (m SecretsModel) Init() tea.Cmd { return nil }

func (m SecretsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m SecretsModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render("ERROR: " + m.err.Error())
	}
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" SECURITY SCAN ") + "\n\n")

	if len(m.matches) == 0 {
		s.WriteString(SecondaryStyle.Render(IconCheck + " No secrets detected in the current directory.") + "\n")
	} else {
		s.WriteString(ErrorStyle.Render(fmt.Sprintf("%s FOUND %d POTENTIAL SECRETS", IconWarning, len(m.matches))) + "\n\n")
		
		for _, match := range m.matches {
			fileInfo := PrimaryStyle.Render(fmt.Sprintf("%s:%d", match.File, match.Line))
			typeInfo := AccentStyle.Render("[" + match.Type + "]")
			
			s.WriteString(fmt.Sprintf("%s %s\n", fileInfo, typeInfo))
			s.WriteString(SubtleStyle.Render("  " + match.Content) + "\n\n")
		}
	}

	s.WriteString(SubtleStyle.Render("Press 'q' to exit"))
	return s.String()
}
