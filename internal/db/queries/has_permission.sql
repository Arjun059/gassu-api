-- name: HasPermission :one
SELECT EXISTS (
    SELECT 1
    FROM users usr
    INNER JOIN role_permissions rp
        ON usr.role_id = rp.role_id
    INNER JOIN permissions p
        ON rp.permission_id = p.id
    WHERE
        usr.id = @user_id
        AND p.resource = @resource
        AND p.action = @action
) AS has_permission;