
-- name: HasPermission :one
SELECT EXISTS (
    SELECT 1
    FROM user_roles ur
    INNER JOIN role_permissions rp
        ON ur.role_id = rp.role_id
    INNER JOIN permissions p
        ON rp.permission_id = p.id
    WHERE
        ur.user_id = $1
        AND p.resource = $2
        AND p.action = $3
) AS has_permission;