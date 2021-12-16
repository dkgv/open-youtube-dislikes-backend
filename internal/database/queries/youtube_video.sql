-- name: UpsertYouTubeVideo :exec
INSERT INTO youtube_video
    (id, likes, dislikes, views, comments, subscribers)
    VALUES ($1, $2, $3, $4, $5, $6)
    ON CONFLICT (id) DO
        UPDATE SET likes = $2, dislikes = $3, views = $4, comments = $5, subscribers = $6;
