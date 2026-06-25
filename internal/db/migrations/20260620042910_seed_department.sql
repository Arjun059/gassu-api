-- +goose Up

INSERT INTO departments (name)
VALUES
    ('Human Resources'),
    ('Engineering'),
    ('Sales'),
    ('Marketing'),
    ('Finance'),
    ('Operations'),
    ('Customer Support')
ON CONFLICT (name) DO NOTHING;

-- +goose Down

DELETE FROM departments
WHERE name IN (
    'Human Resources',
    'Engineering',
    'Sales',
    'Marketing',
    'Finance',
    'Operations',
    'Customer Support'
);