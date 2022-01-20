-- name: InsertComment :exec
INSERT INTO open_youtube_dislikes.comment (video_id, content, negative, neutral, positive, compound) VALUES ($1, $2, $3, $4, $5, $6);
