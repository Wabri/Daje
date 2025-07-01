package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
    and usage of using your command. For example:

    Cobra is a CLI library for Go that empowers applications.
    This application is a tool to generate the needed files
    to quickly create a Cobra application.`,
	Run: runStatusCommand,
}

func runStatusCommand(_ *cobra.Command, _ []string) {
	fmt.Println("status called")
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
