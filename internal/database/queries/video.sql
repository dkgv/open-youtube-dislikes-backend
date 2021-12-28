-- name: FindVideoDetailsByID :one
SELECT * FROM open_youtube_dislikes.video WHERE id = $1;

-- name: FindNVideosByIDHash :many
SELECT * FROM open_youtube_dislikes.video WHERE id_hash LIKE $1 LIMIT $2;

-- name: UpsertVideoDetails :exec
INSERT INTO open_youtube_dislikes.video
    (id, id_hash, likes, dislikes, views, comments, subscribers, published_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    ON CONFLICT (id) DO
        UPDATE SET
            likes = excluded.likes,
            dislikes = excluded.dislikes,
            views = excluded.views,
            comments = excluded.comments,
            subscribers = excluded.subscribers
        WHERE video.likes <= excluded.likes
            AND video.dislikes <= excluded.dislikes
            AND video.views < excluded.views
            AND video.comments <= excluded.comments
            AND video.subscribers <= excluded.subscribers
            AND video.published_at = excluded.published_at;
