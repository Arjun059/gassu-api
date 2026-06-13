-- name: GetPermission :one
SELECT * FROM permissions
WHERE id = $1 LIMIT 1;

-- name: ListPermission :many
SELECT * FROM permissions
ORDER BY created_at;

-- name: CreatePermission :one
INSERT INTO permissions (
  resource, action
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdatePermission :exec
UPDATE permissions
  set resource = $2,
  action = $3
WHERE id = $1;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1;