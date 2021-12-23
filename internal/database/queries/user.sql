-- name: FindUserByID :one
SELECT * FROM "user" WHERE id = $1;

-- name: InsertUser :exec
INSERT INTO "user" (id) VALUES ($1) ON CONFLICT DO NOTHING;
