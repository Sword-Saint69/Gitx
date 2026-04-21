package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type ChatModel struct {
	textInput textinput.Model
	history   []string
	path      string
	quit      bool
}

func NewChatModel(path string) ChatModel {
	ti := textinput.New()
	ti.Placeholder = "Ask me anything about the repo..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 80

	return ChatModel{
		textInput: ti,
		history:   []string{"[i] Welcome to the Gitx Context Chat! How can I help today?"},
		path:      path,
	}
}

func (m ChatModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			query := m.textInput.Value()
			if query == "" {
				return m, nil
			}
			m.history = append(m.history, "> "+query)
			answer, _ := git.AnswerQuestion(m.path, query)
			m.history = append(m.history, "[!] "+answer)
			m.textInput.Reset()
			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quit = true
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m ChatModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" REPO CONTEXT CHAT ") + "\n\n")

	// Show only last 10 messages
	start := 0
	if len(m.history) > 10 {
		start = len(m.history) - 10
	}
	for _, h := range m.history[start:] {
		if strings.HasPrefix(h, ">") {
			s.WriteString(PrimaryStyle.Render(h) + "\n")
		} else {
			s.WriteString(SecondaryStyle.Render(h) + "\n")
		}
	}

	s.WriteString("\n" + m.textInput.View() + "\n\n")
	s.WriteString(SubtleStyle.Render("Press Esc to exit"))

	return s.String()
}
