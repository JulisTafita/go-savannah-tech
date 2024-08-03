package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Mock the TOML configuration file for testing
const mockConfig = `
[server]
run_web_server = true
web_server_host = "localhost"
web_server_port = "8080"

[database]
user_name = "testuser"
user_password = "testpassword"
name = "testdb"
host = "localhost"
port = "3306"

[github]
api_endpoint = "https://api.github.com"
user_token = "testtoken"
repository_name = "testrepo"

[option]
pulling_cron_job_spec = "@every 01h00m00s"
pulling_start_date = "2024-08-03 00:00:00"
pulling_end_date = "2024-08-03 23:59:59"
reset_collection = true
`

func WriteMockConfig(t *testing.T) {
	err := os.WriteFile("./mock-config.toml", []byte(mockConfig), 0644)
	assert.NoError(t, err)
}

func RemoveMockConfig(t *testing.T) {
	err := os.Remove("mock-config.toml")
	assert.NoError(t, err)
}

func TestInitializeConfig(t *testing.T) {
	// Write the mock configuration file
	WriteMockConfig(t)
	defer RemoveMockConfig(t)

	InitializeConfig("./mock-config.toml")

	// Verify that the Default configuration has been initialized correctly
	assert.NotNil(t, Default)
	assert.Equal(t, "localhost", Default.Server.WebServerHost)
	assert.Equal(t, "8080", Default.Server.WebServerPort)
	assert.Equal(t, "testuser", Default.Database.UserName)
	assert.Equal(t, "testpassword", Default.Database.UserPassword)
	assert.Equal(t, "testdb", Default.Database.Name)
	assert.Equal(t, "localhost", Default.Database.Host)
	assert.Equal(t, "3306", Default.Database.Port)
	assert.Equal(t, "https://api.github.com", Default.Github.ApiEndpoint)
	assert.Equal(t, "Bearer testtoken", Default.Github.UserToken) // sanitized
	assert.Equal(t, "testrepo", Default.Github.RepositoryName)
	assert.Equal(t, "@every 01h00m00s", Default.Option.PullingCronJobSpec)
	assert.True(t, Default.Option.ResetCollection)
}

func TestSanitize(t *testing.T) {
	// Write the mock configuration file
	WriteMockConfig(t)
	defer RemoveMockConfig(t)

	InitializeConfig("./mock-config.toml")

	// Verify the sanitization process
	assert.Equal(t, "Bearer testtoken", Default.Github.UserToken)
	assert.Equal(t, "2024-08-03 00:00:00 +0000 UTC", Default.Option.PullingStartDate.String())
	assert.Equal(t, "2024-08-03 23:59:59 +0000 UTC", Default.Option.PullingEndDate.String())
	assert.Equal(t, "@every 01h00m00s", Default.Option.PullingCronJobSpec)
}

func TestGetDatabaseString(t *testing.T) {
	// Write the mock configuration file
	WriteMockConfig(t)
	defer RemoveMockConfig(t)

	InitializeConfig("./mock-config.toml")

	dbString := GetDatabaseString()
	expected := "testuser:testpassword@tcp(localhost:3306)/testdb?parseTime=true"
	assert.Equal(t, expected, dbString)
}

func TestUsePrivateRepository(t *testing.T) {
	// Write the mock configuration file
	WriteMockConfig(t)
	defer RemoveMockConfig(t)

	InitializeConfig("./mock-config.toml")

	assert.True(t, UsePrivateRepository())

	// Test with empty user token
	Default.Github.UserToken = ""
	assert.False(t, UsePrivateRepository())
}
