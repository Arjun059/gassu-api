-- +goose Up

INSERT INTO policies (
    name,
    effect,
    rules
)
VALUES (
    'Super Admin Access',
    'ALLOW',
    jsonb_build_object(
        'offices',
        (
            SELECT jsonb_agg(id)
            FROM offices
        ),
        'departments',
        (
            SELECT jsonb_agg(id)
            FROM departments
        ),
        'hierarchy',
        'ALL'
    )
)
ON CONFLICT (name) DO NOTHING;

-- +goose Down

DELETE FROM policies
WHERE name = 'Super Admin Access';