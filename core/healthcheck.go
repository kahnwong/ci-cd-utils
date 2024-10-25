package core

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"
)

func HealthcheckValidateArgs(args []string) {
	if len(args) == 0 {
		log.Error().Msg("No URL provided")
	} else if len(args) > 1 {
		log.Error().Msg("Too many arguments provided. Only single argument allowed")
	}
}

func validateEndpoint(endpoint string) {
	_, err := url.ParseRequestURI(endpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid URL")
	}
}

func healthcheck(endpoint string) {
	err := requests.
		URL(endpoint).
		Method(http.MethodGet).
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to reach endpoint: %s", endpoint)
	} else {
		log.Info().Msgf("Endpoint healthy: %s", endpoint)
	}
}

func HealthcheckMain(endpoint string) {
	validateEndpoint(endpoint)

	log.Info().Msg("Sleep for 15 seconds")
	time.Sleep(15 * time.Second)

	for range 5 {
		healthcheck(endpoint)
		time.Sleep(250 * time.Millisecond)
	}
}
