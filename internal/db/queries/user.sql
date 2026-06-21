-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT
    u.id,
    u.name,
    u.role_id,
    u.office_id,
    u.department_id,
    u.report_to,
    u.created_at,
    u.updated_at
FROM users u
WHERE EXISTS (
    SELECT 1
        FROM filter_users(
        sqlc.arg(scope),
        sqlc.narg(manager_id),

        sqlc.arg(office_ids)::bigint[],
        sqlc.arg(department_ids)::bigint[],
        sqlc.arg(company_ids)::bigint[],
        sqlc.arg(employment_types)::text[],

        sqlc.narg(my_hierarchy),
        sqlc.narg(hierarchy_mode)
    ) f
    WHERE f.user_id = u.id
)
ORDER BY u.id DESC;

-- name: CreateUser :one
INSERT INTO users (
  name, role_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set name = $2,
  role_id = $3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;