-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsByUser :many
SELECT p.*, f.name AS feed_name FROM posts AS p
    JOIN feeds AS f ON p.feed_id = f.id
    JOIN users AS u ON f.user_id = u.id
WHERE u.id = $1
ORDER BY p.published_at DESC NULLS LAST
LIMIT $2
OFFSET $3;