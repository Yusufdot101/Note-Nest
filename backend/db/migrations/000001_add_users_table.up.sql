CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_updated_at TIMESTAMPTZ,
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL
);

