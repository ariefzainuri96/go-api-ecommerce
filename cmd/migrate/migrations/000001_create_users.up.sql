CREATE TABLE users (
	id BIGSERIAL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE NOT NULL,
	password bytea NOT NULL
)