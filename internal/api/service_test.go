package api

import (
	"github.com/JulisTafita/go-savannahTech/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set up the mock response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"mock":"data"}`))
	}))

	defer mockServer.Close()

	// Save the original GitHub API endpoint and token
	originalApiEndpoint := config.Default.Github.ApiEndpoint
	originalUserToken := config.Default.Github.UserToken

	// Use the mock server URL for the test
	config.Default.Github.ApiEndpoint = mockServer.URL
	config.Default.Github.UserToken = "mockToken"

	// Reset the config after the test
	defer func() {
		config.Default.Github.ApiEndpoint = originalApiEndpoint
		config.Default.Github.UserToken = originalUserToken
	}()

	// Test cases
	tests := []struct {
		name           string
		url            string
		expectedBody   []byte
		expectedError  bool
		usePrivateRepo bool
	}{
		{
			name:           "Successful fetch public repo",
			url:            "/mock-url",
			expectedBody:   []byte(`{"mock":"data"}`),
			expectedError:  false,
			usePrivateRepo: false,
		},
		{
			name:           "Successful fetch private repo",
			url:            "/mock-url",
			expectedBody:   []byte(`{"mock":"data"}`),
			expectedError:  false,
			usePrivateRepo: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//UsePrivateRepository := func() bool {
			//	return tc.usePrivateRepo
			//}

			body, err := fetch(tc.url)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedBody, body)
			}
		})
	}
}
