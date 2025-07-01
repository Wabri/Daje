package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunDaemonCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runDaemonCommand(nil, []string{})
	})

	want := "TODO: daemon calling\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
