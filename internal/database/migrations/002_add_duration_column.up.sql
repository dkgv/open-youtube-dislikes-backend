ALTER TABLE open_youtube_dislikes.video
    ADD COLUMN IF NOT EXISTS duration_sec INTEGER NOT NULL DEFAULT 0;