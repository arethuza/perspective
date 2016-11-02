drop table if exists tenancy;

create table tenancy
(
    id 			        serial      primary key,
    name		        text        unique not null,
    db_password			text	    not null,
    status		        smallint    not null default 0,
    created_at          timestamp   default current_date
);

drop table if exists superuser;

create table superuser
(
    id 			        serial      primary key,
    name		        text        unique not null,
    password_hash	    bytea       not null,
    status		        smallint    not null default 0,
    created_at          timestamp   default current_date
)