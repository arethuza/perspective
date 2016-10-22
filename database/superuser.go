package database

import (
	"database/sql"
	"time"
)

func CreateSuperUser(databaseConnection *sql.DB, username string, password_hash []byte) (int, error) {
	sql := "insert into superuser(name, password_hash, status) " +
		"values($1, $2, 1) returning id"
	var id int
	err := databaseConnection.QueryRow(sql, username, password_hash).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

type SuperUser struct {
	Id           int
	Name         string
	PasswordHash []byte
	Status       int
	CreatedAt    time.Time
}

func ReadSuperUser(databaseConnection *sql.DB, superUserId int) (*SuperUser, error) {
	sql := "select name, password_hash, status, created_at from superuser " +
		"where id = $1"
	var superUser SuperUser
	err := databaseConnection.QueryRow(sql, superUserId).
		Scan(&superUser.Name, &superUser.PasswordHash, &superUser.Status, &superUser.CreatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &superUser, nil
}

func SetSuperUserPasswordHash(databaseConnection *sql.DB, superUserId int, password_hash []byte) error {
	return nil
}