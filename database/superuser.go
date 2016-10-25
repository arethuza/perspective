package database

import (
	"database/sql"
	"fmt"
	"time"
)

const tenancyColumns = "id, name, password_hash, status, created_at"

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
	Id           int	`json:"id"`
	Name         string	`json:"name"`
	PasswordHash []byte	`json:"-"`
	Status       int	`json:"status"`
	CreatedAt    time.Time	`json:"createdAt"`
}

func ReadSuperUserByName(databaseConnection *sql.DB, name string) (*SuperUser, error) {
	sql := fmt.Sprintf("select %s from superuser where name = $1", tenancyColumns)
	return querySingleSuperUser(databaseConnection, sql, name)
}

func ReadSuperUserById(databaseConnection *sql.DB, superUserId int) (*SuperUser, error) {
	sql := fmt.Sprintf("select %s from superuser where id = $1", tenancyColumns)
	return querySingleSuperUser(databaseConnection, sql, superUserId)
}

func querySingleSuperUser(databaseConnection *sql.DB, sql string, value interface{}) (*SuperUser, error) {
	var superUser SuperUser
	err := databaseConnection.QueryRow(sql, value).
		Scan(&superUser.Id, &superUser.Name, &superUser.PasswordHash, &superUser.Status, &superUser.CreatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &superUser, nil
}
