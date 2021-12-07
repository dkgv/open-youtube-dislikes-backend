-- name: AddYouTubeVideo :exec
INSERT INTO youtube_video (content_id, likes, dislikes, views, comment_count) VALUES ($1, $2, $3, $4, $5);
