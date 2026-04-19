package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type reflogItem struct {
	git.ReflogEntry
}

func (i reflogItem) FilterValue() string { return i.Operation + " " + i.Subject }

type undoDelegate struct{}

func (d undoDelegate) Height() int                               { return 1 }
func (d undoDelegate) Spacing() int                              { return 0 }
func (d undoDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d undoDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(reflogItem)
	if !ok {
		return
	}

	hash := SubtleStyle.Render(i.Hash[:7])
	op := AccentStyle.Width(10).Render(i.Operation)
	subj := ValueStyle.Render(i.Subject)

	str := fmt.Sprintf("%s  %s  %s", hash, op, subj)

	fn := func(s ...string) string { return strings.Join(s, " ") }
	if index == m.Index() {
		fn = func(s ...string) string {
			return PrimaryStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type UndoModel struct {
	list     list.Model
	gitPath  string
	err      error
	quitting bool
	done     bool
}

func NewUndoModel(entries []git.ReflogEntry, path string) UndoModel {
	items := make([]list.Item, len(entries))
	for i, e := range entries {
		items[i] = reflogItem{e}
	}

	l := list.New(items, undoDelegate{}, 60, 15)
	l.Title = "Select Point to Restore"
	l.SetShowStatusBar(false)
	l.Styles.Title = TitleStyle

	return UndoModel{
		list:    l,
		gitPath: path,
	}
}

func (m UndoModel) Init() tea.Cmd { return nil }

func (m UndoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			i := m.list.SelectedItem().(reflogItem)
			err := git.UndoTo(m.gitPath, i.Hash)
			if err != nil {
				m.err = err
			} else {
				m.done = true
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m UndoModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}
	if m.done {
		return SecondaryStyle.Render("✅ Repository restored successfully!")
	}
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}
