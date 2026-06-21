-- name: GetResources :one
SELECT * FROM resources
WHERE id = $1 LIMIT 1;

-- name: ListResources :many
SELECT * FROM resources
ORDER BY created_at;

-- name: CreateResources :one
INSERT INTO resources (
  resource, action
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateResources :exec
UPDATE resources
  set resource = $2,
  action = $3
WHERE id = $1;

-- name: DeleteResources :exec
DELETE FROM resources
WHERE id = $1;