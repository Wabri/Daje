package init

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/schroedinger-Hat/Daje/constants"
	"github.com/schroedinger-Hat/Daje/internal/config"
)

func TestSubmitActionReturnsErrorOnInitFailure(t *testing.T) {
	// Ensure daje folder does not exist
	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("unable to get current user: %v", err)
	}
	dajePath := filepath.Join(currentUser.HomeDir, constants.DajeDotFile)
	os.RemoveAll(dajePath)

	expectedErr := errors.New("init failed")
	initEmptyDaje = func() error { return expectedErr }
	defer func() { initEmptyDaje = config.InitEmptyDaje }()

	err = submitAction()
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}
