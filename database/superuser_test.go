package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func TestCreateSuperUser(t *testing.T) {
	databaseConnection, _ := sql.Open("postgres", connectionString)
	name := "superuser" + time.Now().Format(time.RFC3339)
	superUserId, err := CreateSuperUser(databaseConnection, name, []byte("1234"))
	if err != nil {
		t.Error(err)
	}
	if superUserId <= 0 {
		t.Error("invalid superUserId", superUserId)
	}
	superUser, err := ReadSuperUser(databaseConnection, superUserId)
	if err != nil {
		t.Error(err)
	}
	if superUser == nil {
		t.Error("superUser nil")
	}
	if superUser.Name != name {
		t.Error("invalid name", superUser.Name)
	}
	if string(superUser.PasswordHash) != "1234" {
		t.Error("invalid password hash", superUser.PasswordHash)
	}
	if superUser.Status != 1 {
		t.Error("invalid status", superUser.Status)
	}

}
