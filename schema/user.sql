-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
	id serial
		constraint users_pk
			primary key,
	firstname varchar,
	lastname varchar,
	username varchar,
	password varchar,
	email varchar
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;