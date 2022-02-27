-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table users
(
	id int auto_increment,
	firstname varchar(255) null,
	lastname varchar(255) null,
	username varchar(255) null,
	password varchar(255) null,
	email varchar(255) null,
	constraint users_pk
		primary key (id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;