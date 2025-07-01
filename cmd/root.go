/*
Copyright © 2025 Gabriele Puliti <gabriele.puliti@schroedinger-hat.org>

*/
package cmd

import (
	"os"
    "fmt"

	"github.com/spf13/cobra"
)


const Version = "0.1.0"


var rootCmd = &cobra.Command{
	Use:   "daje",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runRootCommand,
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func init() {
}


func runRootCommand(_ *cobra.Command, _ []string) {
  fmt.Println("Daje CLI is running...")
  fmt.Printf("Current version: %s\n", Version)
}
