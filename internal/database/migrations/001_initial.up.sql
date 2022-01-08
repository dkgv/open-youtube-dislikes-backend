CREATE SCHEMA IF NOT EXISTS open_youtube_dislikes;

CREATE TABLE IF NOT EXISTS open_youtube_dislikes.video (
     id TEXT NOT NULL PRIMARY KEY,
     id_hash TEXT NOT NULL,
     likes BIGINT NOT NULL,
     dislikes BIGINT NOT NULL,
     views BIGINT NOT NULL,
     comments BIGINT,
     subscribers BIGINT NOT NULL,
     published_at BIGINT NOT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS video_id_index ON open_youtube_dislikes.video (id);
CREATE INDEX IF NOT EXISTS video_id_hash_index ON open_youtube_dislikes.video (id_hash);
CREATE INDEX IF NOT EXISTS video_updated_at_index ON open_youtube_dislikes.video (updated_at);

CREATE TABLE IF NOT EXISTS open_youtube_dislikes."user" (
    id TEXT NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS user_id_index ON open_youtube_dislikes."user" (id);

CREATE TABLE IF NOT EXISTS open_youtube_dislikes.dislike (
    video_id TEXT NOT NULL REFERENCES video (id),
    user_id TEXT NOT NULL REFERENCES "user" (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS dislike_video_id_index ON open_youtube_dislikes.dislike (video_id);
CREATE INDEX IF NOT EXISTS dislike_user_id_index ON open_youtube_dislikes.dislike (user_id);

CREATE TABLE IF NOT EXISTS open_youtube_dislikes."like" (
     video_id TEXT NOT NULL REFERENCES video (id),
     user_id TEXT NOT NULL REFERENCES "user" (id),
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS like_video_id_index ON open_youtube_dislikes."like" (video_id);
CREATE INDEX IF NOT EXISTS like_user_id_index ON open_youtube_dislikes."like" (user_id);

CREATE TABLE IF NOT EXISTS open_youtube_dislikes.aggregate_dislike (
    id TEXT NOT NULL REFERENCES video (id),
    count INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS video_id_index ON open_youtube_dislikes.aggregate_dislike (id);
