-- +goose Up

CREATE TABLE resources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    resource TEXT NOT NULL,
    action TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (resource, action)
);

CREATE TABLE role_resources (
    role_id UUID NOT NULL,
    resource_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (role_id, resource_id),

    CONSTRAINT fk_role_resources_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role_resources_resource
        FOREIGN KEY (resource_id)
        REFERENCES resources(id)
        ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS role_resources;
DROP TABLE IF EXISTS resources;