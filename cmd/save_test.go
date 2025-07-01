package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunSaveCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runSaveCommand(nil, []string{})
	})

	want := "TODO: save calling\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
