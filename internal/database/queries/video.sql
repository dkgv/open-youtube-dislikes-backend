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
            likes = GREATEST(video.likes, excluded.likes),
            dislikes = GREATEST(video.dislikes, excluded.dislikes),
            views = GREATEST(video.views, excluded.views),
            comments = GREATEST(video.comments, excluded.comments),
            subscribers = GREATEST(video.subscribers, excluded.subscribers)
        WHERE video.likes <= excluded.likes
            OR video.dislikes <= excluded.dislikes
            OR video.views < excluded.views
            OR video.comments <= excluded.comments
            OR video.subscribers <= excluded.subscribers
            OR video.published_at = excluded.published_at;
