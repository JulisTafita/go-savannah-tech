package repository

import (
	"github.com/JulisTafita/go-savannahTech/internal/api"
)

// parseRepository converts an api.Repository object to an IRepository object.
func ParseRepository(apiRepository *api.Repository) *IRepository {
	if apiRepository == nil {
		// Return an empty IRepository if the input is nil.
		return &IRepository{}
	}

	// Map the fields from api.Repository to IRepository.
	return &IRepository{
		RepositoryCreatedAt: apiRepository.CreatedAt,
		RepositoryUpdatedAt: apiRepository.UpdatedAt,
		Name:                apiRepository.Name,
		Owner:               apiRepository.Owner.Login,
		Description:         apiRepository.Description,
		Url:                 apiRepository.HTMLURL,
		Language:            apiRepository.Language,
		ForksCount:          apiRepository.ForksCount,
		StarsCount:          apiRepository.StargazersCount,
		OpenIssuesCount:     apiRepository.OpenIssuesCount,
		WatchersCount:       apiRepository.WatchersCount,
	}
}

// parseCommit converts an api.RepositoryCommit object to an ICommit object.
func parseCommit(apiCommit *api.RepositoryCommit) *ICommit {
	if apiCommit == nil {
		// Return an empty ICommit if the input is nil.
		return &ICommit{}
	}

	// Map the fields from api.RepositoryCommit to ICommit.
	return &ICommit{
		AuthorName:  apiCommit.Commit.Author.Name,
		AuthorEmail: apiCommit.Commit.Author.Email,
		AuthorLogin: apiCommit.Author.Login,
		Date:        apiCommit.Commit.Author.Date,
		Message:     apiCommit.Commit.Message,
		Url:         apiCommit.HTMLURL,
	}
}

// look searches for a repository by name in a slice of api.Repository objects.
func Look(name string, repositories []api.Repository) (repository *IRepository) {
	// Iterate through the repositories to find one with the matching name.
	for i := 0; i < len(repositories); i++ {
		if repositories[i].Name == name {
			// Convert the found repository to an IRepository and return it.
			return ParseRepository(&repositories[i])
		}
	}
	// Return nil if no matching repository is found.
	return nil
}
