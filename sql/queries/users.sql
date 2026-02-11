-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, slug, name, agent_id)
VALUES (
        NOW(),
        NOW(),
        $1,
        $2,
        $3
       )
RETURNING *;

-- name: GetVisitorBySlug :one
SELECT * FROM users WHERE slug = $1;

-- name: SetHasCalled :one
UPDATE users SET has_called = TRUE WHERE slug = $1 RETURNING *;

-- name: IncrementVisit :one
UPDATE users SET visit_count = visit_count + 1 WHERE slug = $1 RETURNING *;