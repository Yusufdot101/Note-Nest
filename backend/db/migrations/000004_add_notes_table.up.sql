CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    likes_count SMALLINT NOT NULL DEFAULT 0,
    comments_count SMALLINT NOT NULL DEFAULT 0,
    color TEXT NOT NULL DEFAULT '#ffffff' CHECK (color ~ '^#[0-9a-fA-F]{6}$'),
    visibility TEXT NOT NULL DEFAULT 'private' CHECK (visibility IN ('private', 'public'))
);

