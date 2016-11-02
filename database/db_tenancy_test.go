package database

import (
	"database/sql"
	"testing"
	"github.com/arethuza/perspective/misc"
)

func TestCreateTenancyDb(t *testing.T) {
	databaseConnection, _ := sql.Open("postgres", connectionString)
	name, _, err := misc.GenerateRandomString(30)
	if err != nil {
		t.Error(err)
	}
	tenancyId, err := CreateTenancy(databaseConnection, name, "1234")
	if err != nil {
		t.Error(err)
	}
	err = CreateTenancyDb(databaseConnection, tenancyId)
	if err != nil {
		t.Error(err)
	}
}
