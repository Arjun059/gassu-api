-- name: HasPermission :one
SELECT
    usr.id,
    usr.name,
    usr.role_id,
    usr.created_at,
    usr.updated_at,

    r.resource_id,
    r.resource,
    r.action,

    CAST(r.resource_id IS NOT NULL AS boolean) AS has_permission
FROM users usr
LEFT JOIN LATERAL (
    SELECT
        p.id AS resource_id,
        p.resource,
        p.action
    FROM role_resources rr
    JOIN resources p
        ON rr.resource_id = p.id
    WHERE
        rr.role_id = usr.role_id
        AND p.resource = sqlc.arg(resource)
        AND p.action = sqlc.arg(action)
    LIMIT 1
) r ON TRUE
WHERE
    usr.id = sqlc.arg(user_id);