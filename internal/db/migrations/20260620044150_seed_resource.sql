-- +goose Up

INSERT INTO resources (resource, action)
VALUES
    -- Users
    ('users', 'create'),
    ('users', 'read'),
    ('users', 'update'),
    ('users', 'delete'),

    -- Roles
    ('roles', 'create'),
    ('roles', 'read'),
    ('roles', 'update'),
    ('roles', 'delete'),

    -- Policies
    ('policies', 'create'),
    ('policies', 'read'),
    ('policies', 'update'),
    ('policies', 'delete'),

    -- Offices
    ('offices', 'create'),
    ('offices', 'read'),
    ('offices', 'update'),
    ('offices', 'delete'),

    -- Departments
    ('departments', 'create'),
    ('departments', 'read'),
    ('departments', 'update'),
    ('departments', 'delete')
ON CONFLICT (resource, action) DO NOTHING;

-- +goose Down

DELETE FROM resources
WHERE (resource, action) IN (
    ('users', 'create'),
    ('users', 'read'),
    ('users', 'update'),
    ('users', 'delete'),

    ('roles', 'create'),
    ('roles', 'read'),
    ('roles', 'update'),
    ('roles', 'delete'),

    ('policies', 'create'),
    ('policies', 'read'),
    ('policies', 'update'),
    ('policies', 'delete'),

    ('offices', 'create'),
    ('offices', 'read'),
    ('offices', 'update'),
    ('offices', 'delete'),

    ('departments', 'create'),
    ('departments', 'read'),
    ('departments', 'update'),
    ('departments', 'delete')
);