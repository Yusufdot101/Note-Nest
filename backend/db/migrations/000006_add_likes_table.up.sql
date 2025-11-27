CREATE TABLE IF NOT EXISTS likes (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_id BIGINT NOT NULL REFERENCES users ON DELETE CASCADE,

    note_id BIGINT,
    comment_id BIGINT,

-- only ONE of them can be set
    CHECK (
        ( note_id IS NOT NULL AND comment_id IS NULL ) OR
        ( note_id IS NULL AND comment_id IS NOT NULL )
    ),

    FOREIGN KEY(note_id) REFERENCES notes(id) ON DELETE CASCADE,
    FOREIGN KEY(comment_id) REFERENCES comments(id) ON DELETE CASCADE

);

-- enforce uniqueness
CREATE UNIQUE INDEX likes_note_unique
ON likes(user_id, note_id)
WHERE note_id IS NOT NULL;

CREATE UNIQUE INDEX likes_comment_unique
ON likes(user_id, comment_id)
WHERE comment_id IS NOT NULL;

