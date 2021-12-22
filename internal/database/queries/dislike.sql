-- name: InsertDislike :exec
INSERT INTO dislike (id, ip_hash) VALUES ($1, $2);

-- name: GetDislikeCount :one
SELECT COUNT(*) AS "count" FROM dislike WHERE id = $1;
