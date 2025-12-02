CREATE TABLE providers (
    id BIGSERIAL PRIMARY KEY,
    provider_name TEXT NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    sub TEXT NOT NULL,
    UNIQUE(provider_name, sub)
);
