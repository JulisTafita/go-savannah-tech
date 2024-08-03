package worker

import (
	"github.com/JulisTafita/go-savannahTech/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func writerResponseByUri(uri string) string {

	if strings.Contains(uri, "/search") {
		return dummyApiRepositorySearchResult
	}

	if strings.Contains(uri, "/commits") {
		if strings.Contains(uri, "&page=1") {
			return "[]"
		}

		return dummyCommitResult
	}

	return ""
}

func TestWorkerOnPullSuccess(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Set up the mock response
		w.Write([]byte(writerResponseByUri(r.RequestURI)))
	}))

	config.Default.Github.ApiEndpoint = mockServer.URL
	config.Default.Github.RepositoryName = "material-design-icons"

	assert.Equal(t, nil, workerJob())
}

func TestWorkerOnDummy4o4(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		// Set up the mock response
		w.Write([]byte(dummy404))
	}))

	config.Default.Github.ApiEndpoint = mockServer.URL
	config.Default.Github.RepositoryName = "material-design-icons"

	assert.Equal(t, nil, workerJob())
}

func TestWorkerOnDummyInternalServerError(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		// Set up the mock response
		w.Write([]byte(dummy404))
	}))

	config.Default.Github.ApiEndpoint = mockServer.URL
	config.Default.Github.RepositoryName = "material-design-icons"

	assert.Equal(t, nil, workerJob())
}
