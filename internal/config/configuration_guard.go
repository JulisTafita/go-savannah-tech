package config

import (
	"fmt"
	"strings"
	"time"
)

// sanitize validates and sanitizes the configuration loaded into Default.
func sanitize() (err error) {
	if Default == nil {
		panic("default configuration is missing") // Panic if the Default configuration is missing.
	}

	if len(strings.TrimSpace(Default.Github.ApiEndpoint)) == 0 {
		panic("github api endpoint is missing") // Panic if the GitHub API endpoint is missing.
	}

	if len(strings.TrimSpace(Default.Github.RepositoryName)) == 0 {
		panic("github repository name is missing") // Panic if the GitHub repository name is missing.
	}

	if len(strings.TrimSpace(Default.Github.UserToken)) == 0 {
		// Print a warning if the GitHub user token is missing, indicating a public repository will be used.
		fmt.Println("⚡ github user token is missing, looking for public repository.")
	} else {
		// Prepend "Bearer " to the GitHub user token.
		Default.Github.UserToken = "Bearer " + strings.TrimSpace(Default.Github.UserToken)
	}

	if Default.Option.PullingStartDate.IsZero() {
		// Print a warning if the pulling start date is not set.
		fmt.Println("⚡ pulling start date not set")
	} else {
		// Print the pulling start date if it is set.
		fmt.Println("💡 Pulling from ", Default.Option.PullingStartDate.Format(time.RFC822))
	}

	// Parse the pulling start date string and set the pulling start date.
	if t, err := time.Parse(time.DateTime, Default.Option.PullingStartDateStr); err == nil {
		Default.Option.PullingStartDate = t
	}

	// Parse the pulling end date string and set the pulling end date.
	if t, err := time.Parse(time.DateTime, Default.Option.PullingEndDateStr); err == nil {
		Default.Option.PullingEndDate = t
	}

	if Default.Option.PullingEndDate.IsZero() {
		// Print a warning if the pulling end date is not set.
		fmt.Println("⚡ pulling end date not set")
	}

	// Check if the pulling cron job spec is set and contains "@every".
	if len(Default.Option.PullingCronJobSpec) > 0 && strings.Contains(Default.Option.PullingCronJobSpec, "@every") {
		// Print the pulling cron job spec.
		fmt.Println("💡 pulling will starts ", Default.Option.PullingCronJobSpec)
	} else {
		// Set a default pulling cron job spec and print it.
		Default.Option.PullingCronJobSpec = "@every 01h00m00s"
		fmt.Println("💡 pulling will starts ", Default.Option.PullingCronJobSpec)
	}

	if Default.Option.ResetCollection {
		// Print a message if the collection will be reset.
		fmt.Println("💡 collection will be reset ")
	}

	if Default.Server.RunWebServer {
		// Set a default web server port if it is not set.
		if len(Default.Server.WebServerPort) == 0 {
			Default.Server.WebServerPort = "8089"
		}
	}

	return err
}
