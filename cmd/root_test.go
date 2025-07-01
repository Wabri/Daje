package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunRootCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runRootCommand(nil, []string{})
	})

	want := "Daje CLI is running...\nCurrent version: " + Version + "\nTODO: return a status of the dotfiles\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
