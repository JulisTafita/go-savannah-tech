package api

import (
	"fmt"
	"strconv"
	"time"
)

type IRepositoryService interface {
	GetRepository(repositoryName string) (Repository, error)
	GetUserRepositoryList() ([]Repository, error)
	SearchRepositoryByName(repositoryName string) ([]Repository, error)
	GetRepositoryCommits(owner, repositoryName string, since, until time.Time) ([]RepositoryCommit, error)
}

type RepositoryService struct{}

func (service *RepositoryService) GetRepository(repositoryName string) (Repository, error) {
	profile, err := fetchProfile()
	if err != nil {
		return Repository{}, err
	}

	return fetchUserRepository(profile.Login, repositoryName)
}

func (service *RepositoryService) GetUserRepositoryList() ([]Repository, error) {
	return fetchUserRepositoryList()
}

func (service *RepositoryService) SearchRepositoryByName(repositoryName string) ([]Repository, error) {
	searchResult, err := fetchSearchResultRepositoryList(repositoryName)
	if err != nil {
		return nil, err
	}
	return searchResult.Items, err
}

func (service *RepositoryService) GetRepositoryCommits(owner, repositoryName string, since, until time.Time) ([]RepositoryCommit, error) {
	var commits []RepositoryCommit
	var newCommits []RepositoryCommit
	var err error
	var activePage int

	for len(newCommits) > 0 || activePage == 0 {
		newCommits = []RepositoryCommit{}
		newCommits, err = fetchRepositoryCommit(owner, repositoryName, since, until, activePage)
		if err != nil {
			fmt.Println(err)
			return commits, nil
		}
		fmt.Println("ℹ️ commit page " + strconv.Itoa(activePage) + " with " + strconv.Itoa(len(newCommits)))
		commits = append(commits, newCommits...)
		activePage += 1
	}

	return commits, nil
}
