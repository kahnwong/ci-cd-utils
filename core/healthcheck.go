package core

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"
)

func HealthcheckValidateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("no URL provided")
	} else if len(args) > 1 {
		return errors.New("too many arguments provided. Only single argument allowed")
	}
	return nil
}

func validateEndpoint(endpoint string) error {
	_, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	return nil
}

func makeRequest(endpoint string) error {
	err := requests.
		URL(endpoint).
		Method(http.MethodGet).
		Fetch(context.Background())

	if err != nil {
		return fmt.Errorf("failed to reach endpoint %s: %w", endpoint, err)
	}

	log.Info().Msgf("Endpoint healthy: %s", endpoint)
	return nil
}

func Healthcheck(endpoint string) error {
	if err := validateEndpoint(endpoint); err != nil {
		return err
	}

	log.Info().Msg("Sleep for 15 seconds")
	time.Sleep(15 * time.Second)

	for range 5 {
		if err := makeRequest(endpoint); err != nil {
			return err
		}
		time.Sleep(250 * time.Millisecond)
	}

	return nil
}
