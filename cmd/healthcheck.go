package cmd

import (
	"github.com/kahnwong/ci-cd-utils/core"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Check whether a URL is reachable",
	Run: func(cmd *cobra.Command, args []string) {
		// go run . healthcheck https://example.com

		if err := core.HealthcheckValidateArgs(args); err != nil {
			log.Fatal().Err(err).Msg("Invalid arguments")
		}

		if err := core.Healthcheck(args[0]); err != nil {
			log.Fatal().Err(err).Msg("Healthcheck failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
}
