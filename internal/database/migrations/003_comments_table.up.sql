CREATE TABLE IF NOT EXISTS open_youtube_dislikes.comment (
    video_id TEXT NOT NULL REFERENCES open_youtube_dislikes.video(id),
    content TEXT NOT NULL,
    negative REAL NOT NULL,
    neutral REAL NOT NULL,
    positive REAL NOT NULL,
    compound REAL NOT NULL
);

CREATE INDEX IF NOT EXISTS comment_video_id_index ON open_youtube_dislikes.comment (video_id);
