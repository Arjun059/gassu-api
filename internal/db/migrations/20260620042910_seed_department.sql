-- +goose Up

INSERT INTO departments (id, name)
VALUES
    (1, 'Human Resources'),
    (2, 'Engineering'),
    (3, 'Sales'),
    (4, 'Marketing'),
    (5, 'Finance'),
    (6, 'Operations'),
    (7, 'Customer Support');

-- Reset sequence
SELECT setval(
    pg_get_serial_sequence('departments', 'id'),
    COALESCE((SELECT MAX(id) FROM departments), 1)
);

-- +goose Down

DELETE FROM departments
WHERE id IN (1, 2, 3, 4, 5, 6, 7);

-- Reset sequence
SELECT setval(
    pg_get_serial_sequence('departments', 'id'),
    COALESCE((SELECT MAX(id) FROM departments), 1)
);