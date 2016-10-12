drop table if exists tenancy;

create table tenancy
(
    id 			        serial      primary key,
    name		        text        unique not null,
    database_name	    text        unique not null,
    admin_user		    text        not null,
    admin_password_hash	text        not null,
    status		        smallint    not null default 0,
    created_create      timestamp   default current_date
)

