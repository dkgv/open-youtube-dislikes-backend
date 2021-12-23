CREATE TABLE IF NOT EXISTS video (
     id TEXT NOT NULL PRIMARY KEY,
     id_hash TEXT NOT NULL,
     likes BIGINT NOT NULL,
     dislikes BIGINT NOT NULL,
     views BIGINT NOT NULL,
     comments BIGINT NOT NULL,
     subscribers BIGINT NOT NULL,
     published_at BIGINT NOT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS video_id_index ON video (id);
CREATE INDEX IF NOT EXISTS video_id_hash_index ON video (id_hash);

CREATE TABLE IF NOT EXISTS "user" (
    id TEXT NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS user_id_index ON "user" (id);

CREATE TABLE IF NOT EXISTS dislike (
    video_id TEXT NOT NULL REFERENCES video (id),
    user_id TEXT NOT NULL REFERENCES "user" (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS dislike_video_id_index ON dislike (video_id);
CREATE INDEX IF NOT EXISTS dislike_user_id_index ON dislike (user_id);

CREATE TABLE IF NOT EXISTS aggregate_dislike (
    id TEXT NOT NULL REFERENCES video (id),
    count INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS video_id_index ON aggregate_dislike (id);
