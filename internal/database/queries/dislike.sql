-- name: InsertDislike :exec
INSERT INTO dislike (id, hashed_ip) VALUES ($1, $2);

-- name: GetDislikeCount :one
SELECT COUNT(*) AS "count" FROM dislike WHERE id = $1;
