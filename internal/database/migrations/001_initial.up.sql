CREATE TABLE IF NOT EXISTS video (
    id TEXT NOT NULL UNIQUE,
    id_hash TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS video_id_index ON video (id);
CREATE INDEX IF NOT EXISTS video_id_hash_index ON video (id);

CREATE TABLE IF NOT EXISTS dislike (
    id TEXT NOT NULL REFERENCES video (id),
    hashed_ip TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS video_id_index ON dislike (id);

CREATE TABLE IF NOT EXISTS youtube_video (
    id TEXT NOT NULL REFERENCES video (id),
    likes BIGINT NOT NULL,
    dislikes BIGINT NOT NULL,
    views BIGINT NOT NULL,
    comments BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS video_id_index ON youtube_video (id);

CREATE TABLE IF NOT EXISTS aggregate_dislike (
    id TEXT NOT NULL REFERENCES video (id),
    count INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS video_id_index ON aggregate_dislike (id);
