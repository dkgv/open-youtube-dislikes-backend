-- name: InsertAggregateDislike :exec
INSERT INTO open_youtube_dislikes.aggregate_dislike (id, count) VALUES ($1, $2);

-- name: UpdateAggregateDislike :exec
UPDATE open_youtube_dislikes.aggregate_dislike SET count = $2, updated_at = $3 WHERE id = $1;

-- name: FindAggregateDislikeByID :one
SELECT count FROM open_youtube_dislikes.aggregate_dislike WHERE id = $1;
