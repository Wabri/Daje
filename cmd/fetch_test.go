package cmd

import (
	"testing"

	"github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunFetchCommand(t *testing.T) {
	got := testutil.CaptureOutput(func() {
		runFetchCommand(nil, []string{})
	})

	want := "TODO: fetch calling\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
