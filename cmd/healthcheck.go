package cmd

import (
	"github.com/kahnwong/ci-cd-utils/core"
	"github.com/spf13/cobra"
)

var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Check whether a URL is reachable",
	Run: func(cmd *cobra.Command, args []string) {
		// go run . healthcheck https://example.com

		core.HealthcheckValidateArgs(args)
		core.HealthcheckMain(args[0])
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
}
