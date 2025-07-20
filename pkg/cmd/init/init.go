package init

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/schroedinger-Hat/Daje/internal/config"
)

// initEmptyDaje is a hook to allow replacing the initialization logic in tests.
var initEmptyDaje = config.InitEmptyDaje

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [flags]",
		Short: "Initialize daje on your system",
		RunE: func(cmd *cobra.Command, args []string) error { //nolint:all
			return submitAction()
		},
	}

	return cmd
}

func submitAction() error {
	if config.IsDajeInitialized() {
		fmt.Println("Daje has been already initialized in the system.")
		return nil
	}

	err := initEmptyDaje()
	if err != nil {
		return err
	}

	fmt.Println("Daje has been initialized successfully!")

	return nil
}
