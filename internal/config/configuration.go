package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"strings"
	"testing"
	"time"
)

// App represents the configuration for the entire application.
type App struct {
	Server   Server   `toml:"server"`
	Database Database `toml:"database"`
	Github   Github   `toml:"github"`
	Option   Option   `toml:"option"`
}

// Database holds the configuration details for the database connection.
type Database struct {
	UserName     string `toml:"user_name"`
	UserPassword string `toml:"user_password"`
	Name         string `toml:"name"`
	Host         string `toml:"host"`
	Port         string `toml:"port"`
}

// Github holds the configuration details for accessing GitHub.
type Github struct {
	ApiEndpoint    string `toml:"api_endpoint"`
	UserToken      string `toml:"user_token"`
	RepositoryName string `toml:"repository_name"`
}

// Server holds the configuration details for the web server.
type Server struct {
	RunWebServer  bool   `toml:"run_web_server"`
	WebServerHost string `toml:"web_server_host"`
	WebServerPort string `toml:"web_server_port"`
}

// Option holds additional configuration options for the application.
type Option struct {
	PullingCronJobSpec  string    `toml:"pulling_cron_job_spec"`
	PullingStartDate    time.Time `toml:"_"` // These fields are populated from their string representations
	PullingEndDate      time.Time `toml:"_"` // and are not directly read from the TOML file.
	PullingStartDateStr string    `toml:"pulling_start_date"`
	PullingEndDateStr   string    `toml:"pulling_end_date"`
	ResetCollection     bool      `toml:"reset_collection"`
}

var Default *App

// init initializes the Default configuration by reading it from a TOML file.
func init() {
	InitializeConfig("./config.toml")
}

func InitializeConfig(config string) {
	Default = &App{}
	_, err := toml.DecodeFile(config, Default)
	if err != nil {
		if testing.Testing() {
			return
		}
		panic(err)
	}

	err = sanitize()
	if err != nil {
		panic(err)
	}
}

// UsePrivateRepository checks if a GitHub user token is provided for private repository access.
func UsePrivateRepository() bool {
	return len(strings.TrimSpace(Default.Github.UserToken)) > 0
}

// GetDatabaseString constructs the database connection string from the configuration.
func GetDatabaseString() string {
	var db = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		Default.Database.UserName,
		Default.Database.UserPassword,
		Default.Database.Host,
		Default.Database.Port,
		Default.Database.Name)

	fmt.Println("mysql : ", db)
	return db
}
