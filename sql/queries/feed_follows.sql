-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES ($1,
            $2,
            $3,
            $4,
            $5)
    RETURNING *
)
SELECT i.*,
       f.name AS feed_name,
       u.name AS user_name
    FROM inserted_feed_follow i
JOIN users u on i.user_id = u.id
JOIN feeds f on i.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT f.name AS feed_name,
       u.name AS user_name
FROM feed_follows i
JOIN users u on i.user_id = u.id
JOIN feeds f on i.feed_id = f.id
WHERE u.name=$1;