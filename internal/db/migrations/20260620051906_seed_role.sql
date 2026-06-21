-- +goose Up

INSERT INTO roles (
    id,
    name,
    identifier,
    hierarchy
)
VALUES
    (
        1,
        'Super Admin',
        'super_admin',
        1
    ),
    (
        2,
        'HR Manager',
        'hr_manager',
        2
    ),
    (
        3,
        'Engineering Manager',
        'engineering_manager',
        2
    ),
    (
        4,
        'Employee',
        'employee',
        3
    );

-- +goose Down

DELETE FROM roles
WHERE id IN (1, 2, 3, 4);
