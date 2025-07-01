package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunSomeCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		// some cli command execution
	})

	want := "Something"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
