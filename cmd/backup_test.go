package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunBackupCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runBackupCommand(nil, []string{})
	})

	want := "TODO: backup calling\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
