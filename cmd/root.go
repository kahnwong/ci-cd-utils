package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "ci-cd-utils",
	Version: version,
	Short:   "Various utils for CI/CD",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
