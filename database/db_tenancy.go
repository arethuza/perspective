package database

import (
	"database/sql"
	"fmt"
)

func CreateTenancyDb(databaseConnection *sql.DB, tenancyId int) (string, error) {
	databaseName := fmt.Sprintf("tenancy_%d", tenancyId)
	sql := fmt.Sprintf("create database %s", databaseName)
	_, err := databaseConnection.Exec(sql)
	if err != nil {
		return "", err
	}
	return databaseName, nil
}

func PopulateTenancyDb(databaseConnection *sql.DB, databaseName string) error {
	_, err := databaseConnection.Exec(tenancyDdl)
	if err != nil {
		return err
	}
	return nil
}

const tenancyDdl = `

drop table if exists user_passwords_hash cascade;
drop table if exists item_version cascade;
drop table if exists item cascade;

create table item
(
	id		serial		primary key,
	name		text    	not null,
	parent_id	integer		references item (id) null,
	current_version	integer		not null
);

create table item_version
(
	item_id		integer		references item (id) not null,
	version		integer		not null,
	item_type	integer		not null,
	created_at  	timestamp   	default current_date,
	created_by	integer		references item (id) not null
);

create table user_passwords_hash
(
	item_id		integer		references item (id) not null,
	password_hash	bytea       	not null
);

alter sequence item_id_seq owned by item.id start 1 increment 1;

insert into item (name, parent_id, current_version) values ('', null, 0);

update item set parent_id=1 where id=1;

alter table item alter parent_id set not null;
`
