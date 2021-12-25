-- name: FindVideoDetailsByID :one
SELECT * FROM open_youtube_dislikes.video WHERE id = $1;

-- name: FindNVideosByIDHash :many
SELECT * FROM open_youtube_dislikes.video WHERE id_hash LIKE $1 LIMIT $2;

-- name: UpsertVideoDetails :exec
INSERT INTO open_youtube_dislikes.video
    (id, id_hash, likes, dislikes, views, comments, subscribers, published_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    ON CONFLICT (id) DO
        UPDATE SET likes = $3, dislikes = $4, views = $5, comments = $6, subscribers = $7
        WHERE likes <= $3 AND dislikes <= $4 AND views < $5 AND comments <= $6 AND subscribers <= $7 AND published_at = $8;
