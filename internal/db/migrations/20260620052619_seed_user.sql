-- +goose Up

INSERT INTO users (
    id,
    name,
    role_id,
    office_id,
    department_id,
    report_to
)
VALUES
(
    1,
    'Super Admin',
    1,
    1,
    1,
    NULL
),
(
    2,
    'HR Manager',
    2,
    1,
    1,
    1
),
(
    3,
    'Engineering Manager',
    3,
    2,
    2,
    1
),
(
    4,
    'HR Executive',
    4,
    1,
    1,
    2
),
(
    5,
    'Software Engineer 1',
    4,
    2,
    2,
    3
),
(
    6,
    'Software Engineer 2',
    4,
    2,
    2,
    3
),
(
    7,
    'Frontend Developer',
    4,
    2,
    2,
    3
),
(
    8,
    'Backend Developer',
    4,
    2,
    2,
    3
),
(
    9,
    'Sales Executive',
    4,
    1,
    3,
    1
),
(
    10,
    'Marketing Executive',
    4,
    1,
    4,
    1
);

-- Reset sequence
SELECT setval(
    pg_get_serial_sequence('users', 'id'),
    COALESCE((SELECT MAX(id) FROM users), 1)
);

-- +goose Down

DELETE FROM users
WHERE id BETWEEN 1 AND 10;

-- Reset sequence
SELECT setval(
    pg_get_serial_sequence('users', 'id'),
    COALESCE((SELECT MAX(id) FROM users), 1)
);