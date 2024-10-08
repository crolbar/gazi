package main

import (
	"fmt"
	"log"
	"os"

	"path/filepath"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)


const defaultWidth = 30
const listHeight = 50

type model struct {
    curr_dir *string;
    saved_selects map[string]string;
    dirs list.Model;
}

func main() {
    p := tea.NewProgram(model(new_model()))
    if _,err := p.Run(); err != nil {
        fmt.Println(err)
    }
}

func new_model() model {
    curr_dir, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }


    curr_dirs := get_dirs(curr_dir);
    dirs := list.New(to_items(curr_dirs), itemDelegate{&curr_dir}, defaultWidth, listHeight)

    dirs.KeyMap.Quit = key.NewBinding(key.WithKeys(""))
    dirs.SetShowHelp(false)

    return model {
        &curr_dir,
        make(map[string]string),
        dirs,
    };
}

func (m *model) go_to_parrent() {
    old_dir := *m.curr_dir

    if old_dir == "/" {
        return;
    }

    m.saved_selects[old_dir] = filepath.Base(m.dirs.SelectedItem().FilterValue())

    //logg(fmt.Sprintf("[SAVIVG]: key: %s, saved_val: %s", old_dir, filepath.Base(m.dirs.SelectedItem().FilterValue())))

    m.go_to_this(get_parrent(*m.curr_dir))
    m.select_this(filepath.Base(old_dir))
}

func (m *model) go_to_cild() {
    selected := m.dirs.SelectedItem().FilterValue()

    var selected_file string
    if *m.curr_dir == "/" {
        selected_file = *m.curr_dir + selected
    } else {
        selected_file = *m.curr_dir + "/" + selected
    }


    if (m.go_to_this(selected_file)) {
        saved, exists := m.saved_selects[*m.curr_dir]

        //logg("[CHECKING SAVED]: " + *m.curr_dir)

        if !exists {
            m.dirs.Select(0)
            return;
        }

        m.select_this(saved)
    }
}

func (m *model) go_to_this(this string) bool {
    if m.dirs.IsFiltered() {
        m.dirs.ResetFilter()
    }

    stat, err := os.Stat(this)
    if err != nil {
        log.Fatal(err)
    }

    if !stat.IsDir() {
        return false;
    }

    old := *m.curr_dir
    *m.curr_dir = this
    dirs := get_dirs(*m.curr_dir)

    if dirs == nil {
        *m.curr_dir = old
        return false;
    }

    m.dirs.SetItems(to_items(dirs))

    return true
}

func (m *model) select_this(this string) {
    items := m.dirs.Items()
    size := len(items)

    //logg("[CHECKING ITEMS]: " + this)
    for i := 0; i < size; i++ {
        //logg(fmt.Sprintf("[FOUND]: idtems: %s, next: %s\n\n", items[i].FilterValue(), this))
        if items[i].FilterValue() == this {
            //logg(fmt.Sprintf("[FOUND]: this: %s\n\n", this))
            m.dirs.Select(i)
            return;
        }
    }

    m.dirs.Select(0)
}

func (m model) Init() tea.Cmd {
    return tea.EnterAltScreen
}
