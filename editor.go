package main

import (
	"os"
	"os/exec"
	"strings"
)

const defaultEditor = "nano"

func getEditor() string {
	if v := os.Getenv("VISUAL"); v != "" {
		return v
	}
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}
	return defaultEditor
}

// editorCmd creates a command to edit the given path using the specified editor string.
func editorCmd(editor, path string) *exec.Cmd {
	if c := strings.Fields(editor); len(c) > 0 {
		return exec.Command(c[0], append(c[1:], path)...)
	}
	return nil
}
