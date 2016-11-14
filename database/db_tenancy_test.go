package database

import (
	"database/sql"
	"fmt"
	"github.com/arethuza/perspective/misc"
	"testing"
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
	databaseName, err := CreateTenancyDb(databaseConnection, tenancyId)
	if err != nil {
		t.Error(err)
	}
	if databaseName != fmt.Sprintf("tenancy_%d", tenancyId) {
		t.Error("invalid databaseName", databaseName)
	}
}

func TestCreatePopulateTenancyDb(t *testing.T) {
	databaseConnection, _ := sql.Open("postgres", connectionString)
	name, _, err := misc.GenerateRandomString(30)
	if err != nil {
		t.Error(err)
	}
	tenancyId, err := CreateTenancy(databaseConnection, name, "1234")
	if err != nil {
		t.Error(err)
	}
	databaseName, err := CreateTenancyDb(databaseConnection, tenancyId)
	err = PopulateTenancyDb(databaseConnection, databaseName)
	if err != nil {
		t.Error(err)
	}
}
