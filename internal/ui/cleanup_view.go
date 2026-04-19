package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitx/internal/git"
)

type branchItem struct {
	name string
	git.Branch
}

func (i branchItem) FilterValue() string { return i.name }

type branchDelegate struct{}

func (d branchDelegate) Height() int                               { return 1 }
func (d branchDelegate) Spacing() int                              { return 0 }
func (d branchDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d branchDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(branchItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.name)

	fn := ValueStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return PrimaryStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type CleanupModel struct {
	list     list.Model
	gitPath  string
	err      error
	quitting bool
}

func NewCleanupModel(branches []git.Branch, path string) CleanupModel {
	items := make([]list.Item, len(branches))
	for i, b := range branches {
		items[i] = branchItem{name: b.Name, Branch: b}
	}

	l := list.New(items, branchDelegate{}, 20, 10)
	l.Title = "Select Branch to Delete"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = TitleStyle

	return CleanupModel{
		list:    l,
		gitPath: path,
	}
}

func (m CleanupModel) Init() tea.Cmd { return nil }

func (m CleanupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			i := m.list.SelectedItem().(branchItem)
			err := git.DeleteBranch(m.gitPath, i.name)
			if err != nil {
				m.err = err
			} else {
				// Refresh list by removing item
				m.list.RemoveItem(m.list.Index())
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m CleanupModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}
