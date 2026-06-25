-- +goose Up

-- Level 1
INSERT INTO users (
    name,
    role_id,
    office_id,
    department_id,
    report_to
)
VALUES (
    'Super Admin',
    (SELECT id FROM roles WHERE identifier = 'super_admin'),
    (SELECT id FROM offices WHERE city = 'Karnal'),
    (SELECT id FROM departments WHERE name = 'Human Resources'),
    NULL
)
ON CONFLICT DO NOTHING;

-- Level 2
INSERT INTO users (
    name,
    role_id,
    office_id,
    department_id,
    report_to
)
VALUES
(
    'HR Manager',
    (SELECT id FROM roles WHERE identifier = 'hr_manager'),
    (SELECT id FROM offices WHERE city = 'Karnal'),
    (SELECT id FROM departments WHERE name = 'Human Resources'),
    (SELECT id FROM users WHERE name = 'Super Admin')
),
(
    'Engineering Manager',
    (SELECT id FROM roles WHERE identifier = 'engineering_manager'),
    (SELECT id FROM offices WHERE city = 'Mohali'),
    (SELECT id FROM departments WHERE name = 'Engineering'),
    (SELECT id FROM users WHERE name = 'Super Admin')
)
ON CONFLICT DO NOTHING;

-- Level 3
INSERT INTO users (
    name,
    role_id,
    office_id,
    department_id,
    report_to
)
VALUES
(
    'HR Executive',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Karnal'),
    (SELECT id FROM departments WHERE name = 'Human Resources'),
    (SELECT id FROM users WHERE name = 'HR Manager')
),
(
    'Software Engineer 1',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Mohali'),
    (SELECT id FROM departments WHERE name = 'Engineering'),
    (SELECT id FROM users WHERE name = 'Engineering Manager')
),
(
    'Software Engineer 2',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Mohali'),
    (SELECT id FROM departments WHERE name = 'Engineering'),
    (SELECT id FROM users WHERE name = 'Engineering Manager')
),
(
    'Frontend Developer',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Mohali'),
    (SELECT id FROM departments WHERE name = 'Engineering'),
    (SELECT id FROM users WHERE name = 'Engineering Manager')
),
(
    'Backend Developer',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Mohali'),
    (SELECT id FROM departments WHERE name = 'Engineering'),
    (SELECT id FROM users WHERE name = 'Engineering Manager')
),
(
    'Sales Executive',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Karnal'),
    (SELECT id FROM departments WHERE name = 'Sales'),
    (SELECT id FROM users WHERE name = 'Super Admin')
),
(
    'Marketing Executive',
    (SELECT id FROM roles WHERE identifier = 'employee'),
    (SELECT id FROM offices WHERE city = 'Karnal'),
    (SELECT id FROM departments WHERE name = 'Marketing'),
    (SELECT id FROM users WHERE name = 'Super Admin')
)
ON CONFLICT DO NOTHING;

-- +goose Down

DELETE FROM users
WHERE name IN (
    'Super Admin',
    'HR Manager',
    'Engineering Manager',
    'HR Executive',
    'Software Engineer 1',
    'Software Engineer 2',
    'Frontend Developer',
    'Backend Developer',
    'Sales Executive',
    'Marketing Executive'
);