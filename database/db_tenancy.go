package database

import (
	"database/sql"
	"fmt"
)

func CreateTenancyDb(databaseConnection *sql.DB, tenancyId int) error {
	sql := fmt.Sprintf("create database tenancy_%d", tenancyId)
	_, err := databaseConnection.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
