package database

import (
	"database/sql"
	"time"
)

func CreateTenancy(databaseConnection *sql.DB, name, password string) (int, error) {
	sql := "insert into tenancy(name, db_password, status) " +
		"values($1, $2, 1) returning id"
	var id int
	err := databaseConnection.QueryRow(sql, name, password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

type Tenancy struct {
	Id        int
	Name      string
	Password  string
	Status    int
	CreatedAt time.Time
}

func ReadTenancy(databaseConnection *sql.DB, tenancyId int) (*Tenancy, error) {
	sql := "select name, db_password, status, created_at from tenancy " +
		"where id = $1"
	var tenancy Tenancy
	err := databaseConnection.QueryRow(sql, tenancyId).
		Scan(&tenancy.Name, &tenancy.Password, &tenancy.Status, &tenancy.CreatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &tenancy, nil
}
