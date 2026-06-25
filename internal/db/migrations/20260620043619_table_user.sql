-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,

    role_id UUID NULL,
    office_id UUID NULL,
    department_id UUID NULL,

    -- Manager/Supervisor of this user
    report_to UUID NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE SET NULL,

    CONSTRAINT fk_user_office
        FOREIGN KEY (office_id)
        REFERENCES offices(id)
        ON DELETE SET NULL,

    CONSTRAINT fk_user_department
        FOREIGN KEY (department_id)
        REFERENCES departments(id)
        ON DELETE SET NULL,

    CONSTRAINT fk_user_report_to
        FOREIGN KEY (report_to)
        REFERENCES users(id)
        ON DELETE SET NULL,

    CONSTRAINT chk_user_report_to
        CHECK (report_to IS NULL OR report_to <> id)
);

-- +goose Down

DROP TABLE IF EXISTS users;