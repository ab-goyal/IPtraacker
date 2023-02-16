package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IP tracker CLI App",
		Long:  `IP tracker CLI App`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
