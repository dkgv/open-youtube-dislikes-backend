-- name: AddYouTubeVideo :exec
INSERT INTO youtube_video (id, likes, dislikes, views, comments) VALUES ($1, $2, $3, $4, $5);
