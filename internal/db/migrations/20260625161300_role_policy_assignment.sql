-- +goose Up

-- Super Admin -> All resources
INSERT INTO role_policy_assignments (
    role_id,
    resource_id,
    policy_id
)
SELECT
    r.id,
    res.id,
    p.id
FROM roles r
CROSS JOIN resources res
JOIN policies p
    ON p.name = 'Super Admin Access'
WHERE r.identifier = 'super_admin'
ON CONFLICT (role_id, resource_id, policy_id) DO NOTHING;


-- HR Manager -> Users CRUD
INSERT INTO role_policy_assignments (
    role_id,
    resource_id,
    policy_id
)
SELECT
    r.id,
    res.id,
    p.id
FROM roles r
JOIN resources res
    ON res.resource = 'users'
JOIN policies p
    ON p.name = 'Karnal HR Lower Access'
WHERE r.identifier = 'hr_manager'
ON CONFLICT (role_id, resource_id, policy_id) DO NOTHING;


-- +goose Down

DELETE FROM role_policy_assignments
WHERE policy_id IN (
    SELECT id
    FROM policies
    WHERE name IN (
        'Super Admin Access',
        'Karnal HR Lower Access'
    )
);