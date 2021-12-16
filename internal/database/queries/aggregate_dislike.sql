-- name: SetAggregateDislikeCount :exec
INSERT INTO aggregate_dislike (id, count) VALUES ($1, $2);

-- name: UpdateAggregateDislikeCount :exec
UPDATE aggregate_dislike SET count = $2, updated_at = $3 WHERE id = $1;

-- name: GetAggregateDislikeCount :one
SELECT count FROM aggregate_dislike WHERE id = $1;
