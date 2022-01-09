-- name: FindVideoDetailsByID :one
SELECT * FROM open_youtube_dislikes.video WHERE id = $1;

-- name: FindNVideosByIDHash :many
SELECT * FROM open_youtube_dislikes.video WHERE id_hash LIKE $1 LIMIT $2;

-- name: UpsertVideoDetails :exec
INSERT INTO open_youtube_dislikes.video
    (id, id_hash, likes, dislikes, views, comments, subscribers, published_at, duration_sec)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    ON CONFLICT (id) DO
        UPDATE SET
            likes = GREATEST(video.likes, excluded.likes),
            dislikes = GREATEST(video.dislikes, excluded.dislikes),
            views = GREATEST(video.views, excluded.views),
            comments = GREATEST(video.comments, excluded.comments),
            subscribers = GREATEST(video.subscribers, excluded.subscribers),
            duration_sec = GREATEST(video.duration_sec, excluded.duration_sec),
            updated_at = NOW()
        WHERE video.likes <= excluded.likes
            OR video.dislikes <= excluded.dislikes
            OR video.views < excluded.views
            OR video.comments <= excluded.comments
            OR video.subscribers <= excluded.subscribers
            OR video.duration_sec <= excluded.duration_sec
            OR video.published_at = excluded.published_at;

-- name: FindNVideosWithoutComments :many
SELECT * FROM open_youtube_dislikes.video WHERE comments <= 0 ORDER BY updated_at LIMIT $1;
