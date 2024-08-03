package worker

import (
	"fmt"
	"github.com/JulisTafita/go-savannahTech/internal/config"
	"github.com/JulisTafita/go-savannahTech/pkg/repository"
	"github.com/robfig/cron"
)

// Run initializes and starts the worker job and sets up a cron job to run periodically.
func Run() (err error) {
	// Run the worker job immediately.
	workerJob()

	// Create a new cron job scheduler.
	job := cron.New()

	// Add the worker job to the cron scheduler with the specified schedule.
	err = job.AddFunc(config.Default.Option.PullingCronJobSpec, workerJob)
	if err != nil {
		return err
	}

	// Start the cron scheduler.
	job.Start()

	// Block forever to keep the program running.
	select {}
}

// workerJob retrieves repository data and processes it.
func workerJob() {
	fmt.Println("🔄 worker start job")

	// Get repository data using the configured repository name.
	err := repository.GetRepositoryData(config.Default.Github.RepositoryName)
	if err != nil {
		fmt.Println("❌ Error getting repository data")
		return
	}

	fmt.Println("✅ job completed")
}
