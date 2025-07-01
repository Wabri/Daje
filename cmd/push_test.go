package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunPushCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runPushCommand(nil, []string{})
	})

	want := "TODO: push calling\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
