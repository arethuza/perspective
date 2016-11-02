package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

var connectionString string = "user=postgres password=password dbname=perspective_manager host=localhost sslmode=disable"

func TestCreateReadTenancy(t *testing.T) {
	databaseConnection, _ := sql.Open("postgres", connectionString)
	name := "tenancy" + time.Now().Format(time.RFC3339)
	tenancyId, err := CreateTenancy(databaseConnection, name, "1234")
	if err != nil {
		t.Error(err)
	}
	if tenancyId <= 0 {
		t.Error("invalid tenancyId", tenancyId)
	}
	tenancy, err := ReadTenancy(databaseConnection, tenancyId)
	if err != nil {
		t.Error(err)
	}
	if tenancy == nil {
		t.Error("tenancy nil")
	}
	if tenancy.Name != name {
		t.Error("invalid name", tenancy.Name)
	}
	if string(tenancy.Password) != "1234" {
		t.Error("invalid password", tenancy.Password)
	}
	if tenancy.Status != 1 {
		t.Error("invalid status", tenancy.Status)
	}
}

func TestReadInvalidTenancy(t *testing.T) {
	databaseConnection, _ := sql.Open("postgres", connectionString)
	tenancy, err := ReadTenancy(databaseConnection, -1)
	if err != nil {
		t.Error(err)
	}
	if tenancy != nil {
		t.Error("read tenancy, nil expected")
	}
}
