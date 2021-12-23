-- name: InsertDislike :exec
INSERT INTO dislike (video_id, user_id) VALUES ($1, $2);

-- name: GetDislikeCount :one
SELECT COUNT(*) AS "count" FROM dislike WHERE video_id = $1;

-- name: DeleteDislike :exec
DELETE FROM dislike WHERE video_id = $1 AND user_id = $2;
