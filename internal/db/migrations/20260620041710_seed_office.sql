-- +goose Up

INSERT INTO offices (
    id,
    address,
    city,
    state,
    pincode
)
VALUES
(
    1,
    'Sector - 3, HSIIDC, Plot 12',
    'Karnal',
    'Haryana',
    '132001'
),
(
    2,
    '316-B, 3rd Floor, Tower-B, Bestech Business Tower, Sector-66',
    'Mohali',
    'Punjab',
    '160062'
);

-- +goose Down

DELETE FROM offices
WHERE id IN (1, 2);