package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Unleash/unleash-client-go/v3"
)

func main() {
	featureToggle, err := unleash.NewClient(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("default"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithRefreshInterval(2*time.Second),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"default:development.unleash-insecure-api-token"}}),
	)
	if err != nil {
		panic(err)
	}

	featureToggle.WaitForReady()

	// With Operational
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		featureEnabled := featureToggle.IsEnabled("toggle-exp")
		if featureEnabled {
			fmt.Fprintf(w, "feature is enabled")
		} else {
			fmt.Fprintf(w, "feature is disabled")
		}
	})

	// With Kill Switch
	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		featureEnabled := featureToggle.IsEnabled("toggle-news")
		if featureEnabled {
			fmt.Fprintf(w, "feature news is enabled")
		} else {
			fmt.Fprintf(w, "feature  news is disabled")
		}
	})

	// With Experiment (with rollout 50% usage)
	http.HandleFunc("/pop-up-banners", func(w http.ResponseWriter, r *http.Request) {
		featureEnabled := featureToggle.IsEnabled("toggle-pop-up-banner")
		if featureEnabled {
			fmt.Fprintf(w, "feature pop-up-banner is enabled")
		} else {
			fmt.Fprintf(w, "feature  pop-up-banner is disabled")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
