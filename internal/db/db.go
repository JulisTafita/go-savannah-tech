package db

import (
	"fmt"
	"github.com/JulisTafita/go-savannahTech/internal/config"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	defaultDriver        = "mysql"     // Database driver
	maxOpenConnexion     = 120         // Max open connections to the database
	maxIdleConnexion     = 8           // Max idle connections to the database
	maxConnexionLifeTime = time.Minute // Max lifetime of a connection
)

var client *sqlx.DB

// init initializes the database connection and sets connection parameters.
func init() {
	var err error
	// Connect to the database using the connection string from the config.
	client, err = sqlx.Connect(defaultDriver, config.GetDatabaseString())
	if err != nil {
		panic(err) // Panic if there's an error connecting to the database.
	}

	// Set database connection parameters.
	client.SetMaxOpenConns(maxOpenConnexion)
	client.SetMaxIdleConns(maxIdleConnexion)
	client.SetConnMaxLifetime(maxConnexionLifeTime)
	client.MapperFunc(strcase.ToSnake) // Use snake case for struct field mapping.
}

// CloseConnexion closes the database connection.
func CloseConnexion() {
	err := client.Close()
	if err != nil {
		fmt.Println("Error closing connection")
		return
	}
}

// Select runs a SELECT query and stores the result in R.
func Select(R any, Q string, A ...any) (err error) {
	err = client.Select(R, Q, A...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// Get runs a SELECT query that is expected to return a single row and stores the result in R.
func Get(R any, Q string, A ...any) (err error) {
	err = client.Get(R, Q, A...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// Exec runs an SQL query that doesn't return rows (e.g., INSERT, UPDATE, DELETE).
func Exec(Q string, A ...any) (err error) {
	_, err = client.Exec(Q, A...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
