-- name: InsertComment :exec
INSERT INTO open_youtube_dislikes.comment (video_id, content, negative, neutral, positive, compound) VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindSentimentByVideoID :one
SELECT AVG(negative) AS negative, AVG(neutral) AS neutral, AVG(positive) AS positive, AVG(compound) AS compound FROM open_youtube_dislikes.comment WHERE video_id = $1;

-- name: FindCommentStatusByVideoID :one
SELECT EXISTS(SELECT 1 FROM open_youtube_dislikes.comment WHERE video_id = $1);