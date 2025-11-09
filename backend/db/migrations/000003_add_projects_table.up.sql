CREATE TABLE IF NOT EXISTS projects (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    visibility TEXT NOT NULL CHECK (visibility IN ('public', 'private')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    color TEXT NOT NULL,
    entries_count SMALLINT NOT NULL DEFAULT 0,
    likes_count SMALLINT NOT NULL DEFAULT 0,
    comments_count SMALLINT NOT NULL DEFAULT 0
);

