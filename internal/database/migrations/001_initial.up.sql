CREATE TABLE IF NOT EXISTS content (
    external_id TEXT NOT NULL UNIQUE,
    external_id_hash TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS external_id_index ON content (external_id);
CREATE INDEX IF NOT EXISTS external_id_hash_index ON content (external_id);

CREATE TABLE IF NOT EXISTS single_dislike (
    content_id TEXT NOT NULL REFERENCES content (external_id),
    hashed_ip TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS content_id_index ON single_dislike (content_id);

CREATE TABLE IF NOT EXISTS aggregate_dislike (
    content_id TEXT NOT NULL REFERENCES content (external_id),
    count INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS content_id_index ON aggregate_dislike (content_id);

CREATE TABLE IF NOT EXISTS youtube_video (
    content_id TEXT NOT NULL REFERENCES content (external_id),
    likes INTEGER NOT NULL,
    dislikes INTEGER NOT NULL,
    views INTEGER NOT NULL,
    comment_count INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS content_id_index ON youtube_video (content_id);