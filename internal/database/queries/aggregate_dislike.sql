-- name: SetAggregateDislikeCount :exec
INSERT INTO aggregate_dislike (content_id, count) VALUES ($1, $2);

-- name: UpdateAggregateDislikeCount :exec
UPDATE aggregate_dislike SET count = $2, updated_at = $3 WHERE content_id = $1;

-- name: GetAggregateDislikeCount :one
SELECT count FROM aggregate_dislike WHERE content_id = $1;
