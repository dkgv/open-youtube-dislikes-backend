-- name: AddDislike :exec
INSERT INTO dislike (content_id, hashed_ip) VALUES ($1, $2);

-- name: GetDislikeCount :one
SELECT COUNT(*) AS "count" FROM dislike WHERE content_id = $1;
