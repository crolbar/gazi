package main

//import (
//	"fmt"
//	"strings"
//)

func (m model) View() string {
    //return fmt.Sprintf(
    //    "curr dir:\n%s", strings.Join(m.dirs, "\n"),
    //)

    return "\n" + m.dirs.View()
}
