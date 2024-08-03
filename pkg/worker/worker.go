package worker

import (
	"fmt"
	"github.com/JulisTafita/go-savannahTech/internal/api"
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
	err = job.AddFunc(config.Default.Option.PullingCronJobSpec, func() {
		errX := workerJob()
		if errX != nil {
			fmt.Println(errX)
			return
		}
	})
	if err != nil {
		return err
	}

	// Start the cron scheduler.
	job.Start()

	// Block forever to keep the program running.
	select {}
}

// workerJob retrieves repository data and processes it.
func workerJob() (err error) {
	fmt.Println("🔄 worker start job")
	var repos = &repository.IRepository{}
	var repositoryName = config.Default.Github.RepositoryName
	var repositoryService api.RepositoryService
	var repositories []api.Repository

	// Check if a private repository should be used.
	if config.UsePrivateRepository() {
		r, _ := repositoryService.GetRepository(repositoryName)
		if r.ID > 0 {
			repos = repository.ParseRepository(&r)
		} else {
			repositories, _ = repositoryService.GetUserRepositoryList()
			repos = repository.Look(repositoryName, repositories)
		}
	}

	if repos != nil && len(repos.Owner) == 0 {
		// Search for the repository by name.
		repositories, _ = repositoryService.SearchRepositoryByName(repositoryName)
		repos = repository.Look(repositoryName, repositories)
	}

	if repos == nil || len(repos.Owner) == 0 {
		fmt.Println("❌ Error getting repository data")
		return
	}

	// Add or update the repository in the database.
	err = repos.AddOrUpdate()
	if err != nil {
		fmt.Println("❌ Error getting repository data")
		return
	}

	fmt.Println("✅ repository name : ", repos.Name, " by ", repos.Owner)

	// Sync repository commits.
	err = repos.SyncRepositoryCommits()
	if err != nil {
		fmt.Println("❌ Error getting repository data")
		return
	}

	fmt.Println("✅ job completed")
	return nil
}
