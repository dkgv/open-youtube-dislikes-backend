-- name: InsertDislike :exec
INSERT INTO dislike (id, ip_hash) VALUES ($1, $2);

-- name: GetDislikeCount :one
SELECT COUNT(*) AS "count" FROM dislike WHERE id = $1;

-- name: DeleteDislike :exec
DELETE FROM dislike WHERE id = $1 AND ip_hash = $2;
