package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type GenMsgModel struct {
	suggestion string
	quit       bool
}

func NewGenMsgModel(s string) GenMsgModel {
	return GenMsgModel{suggestion: s}
}

func (m GenMsgModel) Init() tea.Cmd { return nil }

func (m GenMsgModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
		if msg.String() == "y" {
			// In a real app, this would trigger the actual commit
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m GenMsgModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" AI COMMIT ASSISTANT ") + "\n\n")
	s.WriteString(SubtleStyle.Render("I analyzed your staged diff and suggest this message:") + "\n\n")
	
	s.WriteString(BlockHeader(" Suggested Message ") + "\n")
	s.WriteString("  " + ValueStyle.Render(m.suggestion) + "\n\n")

	s.WriteString(SubtleStyle.Render("Press 'y' to accept (copies to clipboard) or 'q' to cancel."))
	return s.String()
}
