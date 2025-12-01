CREATE TABLE providers (
    id BIGSERIAL PRIMARY KEY,
    providers_name TEXT NOT NULL,
    user_sub BIGINT NOT NULL REFERENCES users ON DELETE CASCADE,
    UNIQUE(providers_name, user_sub)
);
