package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunApplyCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runApplyCommand(nil, []string{})
	})

	want := "apply called\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
