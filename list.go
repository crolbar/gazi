package main


import (
	"fmt"
	"io"
	//"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
type item string

type itemDelegate struct {
    curr_dir *string;
}

func (i item) FilterValue() string { return string(i) }

func to_items(strs []string) []list.Item {
    items := []list.Item{}

    for i := 0; i < len(strs); i++ {
        items = append(items, item(strs[i]))
    }

    return items;
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

    //logg("[ITEM]: " + string(i))
    var path string
    if *d.curr_dir == "/" {
        path = *d.curr_dir + string(i)
    } else {
        path = *d.curr_dir + "/" + string(i)
    }

    //logg("[CURR]: " + path)
    fmt_item, fn := format_item(string(i), path, itemStyle)


	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(fmt_item))
}
