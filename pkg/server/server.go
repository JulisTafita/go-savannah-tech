package server

import (
	"fmt"
	"github.com/JulisTafita/go-savannahTech/internal/config"
	"github.com/JulisTafita/go-savannahTech/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Serve starts the web server if configured to do so.
func Serve() (err error) {
	// If the server is not configured to run, block forever.
	if !config.Default.Server.RunWebServer {
		select {}
	}

	// Set the Gin framework to release mode.
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Define the route for getting commit authors.
	// Example: http://localhost:8089/commit-author?number_of_author=1
	router.GET("/commit-author", getCommitAuthor)

	// Define the route for getting repository commits.
	// Example: http://localhost:8089/repository-commit?repository=abeg-rebuild
	router.GET("/repository-commit", getRepositoryCommits)

	// Build the server address from the configured host and port.
	var serverHost = config.Default.Server.WebServerHost + ":" + config.Default.Server.WebServerPort

	// Print the server address to the console.
	fmt.Println("💡 Listening on " + serverHost)

	// Run the Gin router on the specified address.
	err = router.Run(serverHost)
	if err != nil {
		return err
	}

	return err
}

// getRepositoryCommits handles the /repository-commit route.
func getRepositoryCommits(c *gin.Context) {
	var repositoryName string
	var repos repository.IRepository
	var err error

	// Get the repository name from the query parameters.
	repositoryName = c.Query("repository")
	if repositoryName == "" {
		// Return a 400 Bad Request response if the repository name is missing.
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "repository name is required",
		})
		return
	}

	// Retrieve the repository data from the database.
	repos, err = repository.GetRepository()
	if err != nil {
		// Return a 500 Internal Server Error response if something goes wrong.
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong while trying to query repository commits",
		})
		return
	}

	// Return the repository commits as a JSON response.
	c.JSON(http.StatusOK, gin.H{
		"data": repos.Commits,
	})
}

// getCommitAuthor handles the /commit-author route.
func getCommitAuthor(c *gin.Context) {
	var N int
	var err error

	// Get the number of authors from the query parameters and convert it to an integer.
	N, err = strconv.Atoi(c.Query("number_of_author"))
	if err != nil {
		// Return a 500 Internal Server Error response if something goes wrong.
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong while trying to query commit author",
		})
		return
	}

	// Retrieve the top N commit authors from the database.
	commitAuthor, err := repository.GetTopNCommitAuthors(N)
	if err != nil {
		// Return a 500 Internal Server Error response if something goes wrong.
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong while trying to query commit author",
		})
		return
	}

	// Return the commit authors as a JSON response.
	c.JSON(http.StatusOK, gin.H{
		"data": commitAuthor,
	})
}
