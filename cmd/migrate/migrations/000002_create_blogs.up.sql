create table if not exists blogs (
    id bigserial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);