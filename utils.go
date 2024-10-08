package main

import (
	"log"
	"os"
	"path/filepath"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func get_dirs(wd string) []string {
    var dirs []string


    files, err := os.ReadDir(wd)
    if err != nil {
        if os.IsPermission(err) {
            return nil
        }

        log.Fatal(err)
    }


    for _, file := range files {
        dirs = append(dirs, file.Name())
    }

    return dirs;
}

func get_parrent(wd string) string {
    return filepath.Join(wd, "..")
}

func logg(s string) {
    file, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString(s + "\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
}


func format_item(i, path string, style lipgloss.Style) (string, func(strs ...string) string) {
    stat, err := os.Stat(path)
    if err == nil {
        //logg("error: " + err.Error())
        //log.Fatal(err)
        if (stat.IsDir()) {
            return "󰉋 " + i, style.Foreground(lipgloss.Color("4")).Render;
        }
    }


    switch filepath.Ext(i) {
    case ".nix":
        return " " + i, style.Foreground(lipgloss.Color("110")).Render;
    case ".go":
        return " " + i, style.Foreground(lipgloss.Color("74")).Render;
    case ".rs":
        return " " + i, style.Foreground(lipgloss.Color("173")).Render;
    case ".c":
        return " " + i, style.Foreground(lipgloss.Color("31")).Render;
    case ".h":
        return " " + i, style.Foreground(lipgloss.Color("97")).Render;
    case ".java":
        return " " + i, style.Foreground(lipgloss.Color("131")).Render;
    case ".lock":
        return " " + i, style.Foreground(lipgloss.Color("243")).Render;
    }

    switch i {
    case "go.mod", "go.sum":
        return " " + i, style.Foreground(lipgloss.Color("74")).Render;
    case ".gitignore", ".gitmodules":
        return " " + i, style.Render;
    }

    return "󰈔 " + i , style.Render;
}
