package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetEditor(t *testing.T) {
	tt := []struct {
		name     string
		visual   string
		editor   string
		expected string
	}{
		{
			name:     "default",
			expected: "nano",
		},
		{
			name:     "$EDITOR only",
			editor:   "vim",
			expected: "vim",
		},
		{
			name:     "$VISUAL only",
			visual:   "code -w",
			expected: "code -w",
		},
		{
			name:     "both set - $VISIAL wins",
			visual:   "code -w",
			editor:   "vim",
			expected: "code -w",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.visual != "" {
				os.Setenv("VISUAL", tc.visual)
			} else {
				os.Unsetenv("VISUAL")
			}
			if tc.editor != "" {
				os.Setenv("EDITOR", tc.editor)
			} else {
				os.Unsetenv("EDITOR")
			}

			got := getEditor()
			if got != tc.expected {
				t.Errorf("getEditor() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestEditorCmd(t *testing.T) {
	tt := []struct {
		name     string
		editor   string
		path     string
		wantCmd  string
		wantArgs []string
	}{
		{
			name:     "simple editor",
			editor:   "nano",
			path:     "test.txt",
			wantCmd:  "nano",
			wantArgs: []string{"test.txt"},
		},
		{
			name:     "editor with flags",
			editor:   "code --wait",
			path:     "test.txt",
			wantCmd:  "code",
			wantArgs: []string{"--wait", "test.txt"},
		},
		{
			name:     "empty editor",
			editor:   "",
			path:     "test.txt",
			wantCmd:  "",
			wantArgs: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cmd := editorCmd(tc.editor, tc.path)
			if tc.editor == "" {
				if cmd != nil {
					t.Errorf("editorCmd() = %v, want nil", cmd)
				}
				return
			}

			if filepath.Base(cmd.Path) != tc.wantCmd {
				t.Errorf("cmd.Path = %v, want %v", cmd.Path, tc.wantCmd)
			}

			if len(cmd.Args) < 2 {
				t.Errorf("cmd.Args too short = %v", cmd.Args)
				return
			}

			gotArgs := cmd.Args[1:]
			if len(gotArgs) != len(tc.wantArgs) {
				t.Errorf("cmd.Args = %v, want %v", gotArgs, tc.wantArgs)
			}

			for i := range gotArgs {
				if gotArgs[i] != tc.wantArgs[i] {
					t.Errorf("cmd.Args[%d] = %v, want %v", i, gotArgs[i], tc.wantArgs[i])
				}
			}
		})
	}
}
