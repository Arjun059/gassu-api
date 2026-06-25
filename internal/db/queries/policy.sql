-- name: GetPolicies :many

SELECT
    p.id,
    p.name,
    p.effect,
    p.rules
FROM user_policy_assignments upa
JOIN policies p
    ON p.id = upa.policy_id
JOIN resources r
    ON r.id = upa.resource_id
WHERE
    upa.user_id = sqlc.arg(user_id)
    AND r.resource = sqlc.arg(resource)
    AND r.action = sqlc.arg(action)

UNION ALL

SELECT
    p.id,
    p.name,
    p.effect,
    p.rules
FROM users u
JOIN role_policy_assignments rpa
    ON rpa.role_id = u.role_id
JOIN policies p
    ON p.id = rpa.policy_id
JOIN resources r
    ON r.id = rpa.resource_id
WHERE
    u.id = sqlc.arg(user_id)
    AND r.resource = sqlc.arg(resource)
    AND r.action = sqlc.arg(action);