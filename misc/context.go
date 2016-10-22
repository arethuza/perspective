package misc

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
)

type Context struct {
	Config             *Config
	DatabaseConnection *sql.DB
}

func CreateContext(path string, config *Config) (*Context, error) {
	context := &Context{}
	a := strings.Split(path, "/")
	var connectionString string
	switch len(a) {
	case 1, 2:
		// Accessing a tenancy so use the database connection from the config
		connectionString = config.DbConnection
	default:
		// Use the database for the tenancy
		connectionString = ""
	}
	var err error
	context.DatabaseConnection, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	context.Config = config
	return context, nil
}
