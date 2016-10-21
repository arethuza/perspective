package database

import (
	"database/sql"
	"time"
)

func CreateTenancy(databaseConnection *sql.DB, name, username string, password_hash []byte) (int, error) {
	sql := "insert into tenancy(name, admin_user, admin_password_hash, status) " +
		"values($1, $2, $3, 1) returning id"
	var id int
	err := databaseConnection.QueryRow(sql, name, username, password_hash).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

type Tenancy struct {
	Id           int
	Name         string
	Username     string
	PasswordHash []byte
	Status       int
	CreatedAt    time.Time
}

func ReadTenancy(databaseConnection *sql.DB, tenancyId int) (*Tenancy, error) {
	sql := "select name, admin_user, admin_password_hash, status, created_at from tenancy " +
		"where id = $1"
	var tenancy Tenancy
	err := databaseConnection.QueryRow(sql, tenancyId).
		Scan(&tenancy.Name, &tenancy.Username, &tenancy.PasswordHash, &tenancy.Status, &tenancy.CreatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &tenancy, nil
}
