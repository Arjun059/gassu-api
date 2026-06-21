-- name: GetPolicies :many

SELECT
    p.effect,
    p.rules
FROM user_policy_assignments upa
JOIN policies p
    ON p.id = upa.policy_id
WHERE
    upa.user_id = sqlc.arg(user_id)
    AND upa.resource_id = sqlc.arg(resource_id)

UNION ALL

SELECT
    p.effect,
    p.rules
FROM users u
JOIN role_policy_assignments rpa
    ON rpa.role_id = u.role_id
JOIN policies p
    ON p.id = rpa.policy_id
WHERE
    u.id = sqlc.arg(user_id)
    AND rpa.resource_id = sqlc.arg(resource_id);


-- name: GetUserHierarchy :one

SELECT
    r.hierarchy
FROM users u
JOIN roles r
    ON r.id = u.role_id
WHERE
    u.id = sqlc.arg(user_id);