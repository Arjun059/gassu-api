-- +goose Up

INSERT INTO offices (
    address,
    city,
    state,
    pincode
)
VALUES
(
    'Sector - 3, HSIIDC, Plot 12',
    'Karnal',
    'Haryana',
    '132001'
),
(
    '316-B, 3rd Floor, Tower-B, Bestech Business Tower, Sector-66',
    'Mohali',
    'Punjab',
    '160062'
);

-- +goose Down

DELETE FROM offices
WHERE (address, city, state, pincode) IN (
    (
        'Sector - 3, HSIIDC, Plot 12',
        'Karnal',
        'Haryana',
        '132001'
    ),
    (
        '316-B, 3rd Floor, Tower-B, Bestech Business Tower, Sector-66',
        'Mohali',
        'Punjab',
        '160062'
    )
);