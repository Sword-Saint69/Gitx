package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type PulseMsg time.Time

type PulseModel struct {
	path      string
	activity  []string
	pulseLine string
	tick      int
	quit      bool
}

func NewPulseModel(path string) PulseModel {
	return PulseModel{
		path:      path,
		pulseLine: "---^---",
	}
}

func tickPulse() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return PulseMsg(t)
	})
}

func (m PulseModel) Init() tea.Cmd {
	return tickPulse()
}

func (m PulseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	case PulseMsg:
		m.tick++
		// Rotate pulse line
		pulses := []string{"---^---", "--^----", "-^-----", "^------", "------^", "-----^-", "----^--"}
		m.pulseLine = pulses[m.tick%len(pulses)]
		
		// Get fresh activity
		standup, _ := git.GetStandup(m.path)
		if len(standup) > 0 {
			m.activity = standup[0].Commits
		}
		
		return m, tickPulse()
	}
	return m, nil
}

func (m PulseModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(HeaderStyle.Render(" REPOSITORY PULSE ") + "\n\n")
	
	s.WriteString(AccentStyle.Render("Heartbeat: ") + fmt.Sprintf("[%s]", ValueStyle.Render(m.pulseLine)) + "\n\n")

	s.WriteString(PrimaryStyle.Render("RECENT MOVEMENT") + "\n")
	if len(m.activity) == 0 {
		s.WriteString(SubtleStyle.Render("  No activity detected in the last 24h.") + "\n")
	} else {
		for i, a := range m.activity {
			if i > 5 { break }
			s.WriteString(fmt.Sprintf("  %s %s\n", SecondaryStyle.Render(">"), a))
		}
	}

	s.WriteString("\n" + SubtleStyle.Render("Watching for changes... Press 'q' to exit"))
	return s.String()
}
