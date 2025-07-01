package cmd

import (
	"testing"

    "github.com/schrodinger-hat/daje/internal/testutil"
)

func TestRunstatusCommand(t *testing.T) {
  got := testutil.CaptureOutput(func() {
    runStatusCommand(nil, []string{})
  })

  want := "status called\n"

  if got != want {
    t.Errorf("Expected %q, got %q", want, got)
  }
}
