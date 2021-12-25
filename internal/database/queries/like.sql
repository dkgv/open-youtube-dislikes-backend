-- name: InsertLike :exec
INSERT INTO open_youtube_dislikes.like (video_id, user_id) VALUES ($1, $2);

-- name: GetLikeCount :one
SELECT COUNT(*) AS "count" FROM open_youtube_dislikes.like WHERE video_id = $1;

-- name: DeleteLike :exec
DELETE FROM open_youtube_dislikes.like WHERE video_id = $1 AND user_id = $2;

-- name: FindLike :one
SELECT * FROM open_youtube_dislikes.like WHERE video_id = $1 AND user_id = $2;