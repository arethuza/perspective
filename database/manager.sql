drop table if exists tenancy;

create table tenancy
(
    id 			        serial      primary key,
    name		        text        unique not null,
    admin_user		    text        not null,
    admin_password_hash	bytea        not null,
    status		        smallint    not null default 0,
    created_at      timestamp   default current_date
)

