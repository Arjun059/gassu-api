-- +goose Up

INSERT INTO resources (id, resource, action)
VALUES
    -- Users
    (1, 'users', 'create'),
    (2, 'users', 'read'),
    (3, 'users', 'update'),
    (4, 'users', 'delete'),

    -- Roles
    (5, 'roles', 'create'),
    (6, 'roles', 'read'),
    (7, 'roles', 'update'),
    (8, 'roles', 'delete'),

    -- Policies
    (9, 'policies', 'create'),
    (10, 'policies', 'read'),
    (11, 'policies', 'update'),
    (12, 'policies', 'delete'),

    -- Offices
    (13, 'offices', 'create'),
    (14, 'offices', 'read'),
    (15, 'offices', 'update'),
    (16, 'offices', 'delete'),

    -- Departments
    (17, 'departments', 'create'),
    (18, 'departments', 'read'),
    (19, 'departments', 'update'),
    (20, 'departments', 'delete');


-- +goose Down

DELETE FROM resources
WHERE id BETWEEN 1 AND 20;
