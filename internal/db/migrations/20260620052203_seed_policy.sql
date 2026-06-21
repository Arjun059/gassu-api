-- +goose Up

INSERT INTO policies (
    id,
    name,
    effect,
    rules
)
VALUES
(
    1,
    'Karnal Office Access',
    'ALLOW',
    '{
        "offices": [1]
    }'::jsonb
),
(
    2,
    'Mohali Office Access',
    'ALLOW',
    '{
        "offices": [2]
    }'::jsonb
),
(
    3,
    'HR Department Access',
    'ALLOW',
    '{
        "departments": [1]
    }'::jsonb
),
(
    4,
    'Engineering Department Access',
    'ALLOW',
    '{
        "departments": [2]
    }'::jsonb
),
(
    5,
    'Lower Hierarchy Access',
    'ALLOW',
    '{
        "hierarchy": "LOWER"
    }'::jsonb
);

-- +goose Down

DELETE FROM policies
WHERE id IN (1, 2, 3, 4, 5);
