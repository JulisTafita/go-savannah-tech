package mock

import (
	"github.com/JulisTafita/go-savannahTech/internal/api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockRepositoryService struct{}

var mockRepository = api.Repository{}
var mockSliceRepository = []api.Repository{}
var mockApiRepositoryCommit = []api.RepositoryCommit{}

func (service *MockRepositoryService) GetRepository(repositoryName string) (api.Repository, error) {
	return mockRepository, nil
}

func (service *MockRepositoryService) GetUserRepositoryList() ([]api.Repository, error) {
	return mockSliceRepository, nil
}

func (service *MockRepositoryService) SearchRepositoryByName(repositoryName string) ([]api.Repository, error) {
	return mockSliceRepository, nil
}

func (service *MockRepositoryService) GetRepositoryCommits(owner, repositoryName string, since, until time.Time) ([]api.RepositoryCommit, error) {
	return mockApiRepositoryCommit, nil
}

func TestInitRepositoryService(t *testing.T) {}

func TestGetRepository(t *testing.T) {
	var repository MockRepositoryService
	type args struct {
		repositoryName string
	}

	tests := []struct {
		repositoryName string
		want           api.Repository
	}{
		{repositoryName: "", want: mockRepository},
		{repositoryName: "nativel", want: mockRepository},
		{repositoryName: "material-design-icons", want: mockRepository},
	}

	for _, tt := range tests {
		t.Run(tt.repositoryName, func(t *testing.T) {
			got, _ := repository.GetRepository(tt.repositoryName)
			assert.Equal(t, tt.want, got)
		})
	}
}
