-- +goose Up

INSERT INTO roles (
    name,
    identifier,
    hierarchy
)
VALUES
    (
        'Super Admin',
        'super_admin',
        1
    ),
    (
        'HR Manager',
        'hr_manager',
        2
    ),
    (
        'Engineering Manager',
        'engineering_manager',
        2
    ),
    (
        'Employee',
        'employee',
        3
    )
ON CONFLICT (identifier) DO NOTHING;

-- +goose Down

DELETE FROM roles
WHERE identifier IN (
    'super_admin',
    'hr_manager',
    'engineering_manager',
    'employee'
);