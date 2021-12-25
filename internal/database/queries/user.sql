-- name: FindUserByID :one
SELECT * FROM open_youtube_dislikes."user" WHERE id = $1;

-- name: InsertUser :exec
INSERT INTO open_youtube_dislikes."user" (id) VALUES ($1) ON CONFLICT DO NOTHING;
