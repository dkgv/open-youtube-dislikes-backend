-- name: InsertDislike :exec
INSERT INTO open_youtube_dislikes.dislike (video_id, user_id) VALUES ($1, $2);

-- name: GetDislikeCount :one
SELECT COUNT(*) AS "count" FROM open_youtube_dislikes.dislike WHERE video_id = $1;

-- name: DeleteDislike :exec
DELETE FROM open_youtube_dislikes.dislike WHERE video_id = $1 AND user_id = $2;

-- name: FindDislike :one
SELECT * FROM open_youtube_dislikes.dislike WHERE video_id = $1 AND user_id = $2;