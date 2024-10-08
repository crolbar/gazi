package main


import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.dirs.SetWidth(msg.Width)
        m.dirs.SetHeight(msg.Height)
        return m, nil

    case tea.KeyMsg:

        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        }
    }

    if m.dirs.FilterState() != list.Filtering {
        mo, c := m.char_update(msg)

        if mo != nil {
            return mo, c
        }
    }

	var cmd tea.Cmd
	m.dirs, cmd = m.dirs.Update(msg)
    return m, cmd
}

func (m model) char_update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit

        case "h":
            m.go_to_parrent()
            return m, nil
        case "l":
            m.go_to_cild()
            return m, nil
        }
    }

    return nil, nil
}
