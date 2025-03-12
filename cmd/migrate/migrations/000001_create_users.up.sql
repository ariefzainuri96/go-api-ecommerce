CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);