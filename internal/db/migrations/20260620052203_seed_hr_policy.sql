-- +goose Up

INSERT INTO policies (
    name,
    effect,
    rules
)
VALUES (
    'Karnal HR Lower Access',
    'ALLOW',
    jsonb_build_object(
        'offices',
        jsonb_build_array(
            (SELECT id FROM offices WHERE city = 'Karnal')
        ),
        'departments',
        jsonb_build_array(
            (SELECT id FROM departments WHERE name = 'Human Resources')
        ),
        'hierarchy',
        'LOWER'
    )
)
ON CONFLICT (name) DO NOTHING;

-- +goose Down

DELETE FROM policies
WHERE name = 'Karnal HR Lower Access';