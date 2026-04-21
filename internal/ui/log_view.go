package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type logItem struct {
	git.LogEntry
}

func (i logItem) FilterValue() string { return i.Subject + " " + i.Author }

type logDelegate struct{}

func (d logDelegate) Height() int                               { return 1 }
func (d logDelegate) Spacing() int                              { return 0 }
func (d logDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d logDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(logItem)
	if !ok {
		return
	}

	hash := AccentStyle.Render(i.Hash[:7])
	date := SubtleStyle.Render(i.Date)
	subj := ValueStyle.Render(i.Subject)

	str := fmt.Sprintf("%s %s %s", hash, date, subj)

	fn := func(s ...string) string { return strings.Join(s, " ") }
	if index == m.Index() {
		fn = func(s ...string) string {
			return PrimaryStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type LogModel struct {
	list     list.Model
	quit     bool
}

func NewLogModel(entries []git.LogEntry) LogModel {
	items := make([]list.Item, len(entries))
	for i, e := range entries {
		items[i] = logItem{e}
	}

	l := list.New(items, logDelegate{}, 80, 20)
	l.Title = "Commit History"
	l.SetShowStatusBar(false)
	l.Styles.Title = TitleStyle

	return LogModel{list: l}
}

func (m LogModel) Init() tea.Cmd { return nil }

func (m LogModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m LogModel) View() string {
	if m.quit {
		return ""
	}
	return "\n" + m.list.View()
}
