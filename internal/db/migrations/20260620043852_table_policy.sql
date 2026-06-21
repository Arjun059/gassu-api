-- +goose Up

CREATE TYPE policy_effect AS ENUM (
    'ALLOW',
    'DENY'
);

CREATE TABLE policies (
    id BIGSERIAL PRIMARY KEY,

    name TEXT NOT NULL UNIQUE,
    effect policy_effect NOT NULL DEFAULT 'ALLOW',

    -- Example:
    -- {
    --   "offices": [1,2],
    --   "departments": [3],
    --   "companies": [1],
    --   "hierarchy": "LOWER"
    -- }
    rules JSONB NOT NULL DEFAULT '{}'::jsonb,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_policies_rules_object
    CHECK (jsonb_typeof(rules) = 'object')
);

CREATE INDEX idx_policies_rules_gin
ON policies
USING GIN (rules);

CREATE TABLE role_policy_assignments (
    id BIGSERIAL PRIMARY KEY,

    role_id BIGINT NOT NULL,
    resource_id BIGINT NOT NULL,
    policy_id BIGINT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_role_resource_policy
        UNIQUE (role_id, resource_id, policy_id),

    CONSTRAINT fk_rpa_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_rpa_permission
        FOREIGN KEY (resource_id)
        REFERENCES resources(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_rpa_policy
        FOREIGN KEY (policy_id)
        REFERENCES policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_rpa_lookup
ON role_policy_assignments(role_id, resource_id);

CREATE INDEX idx_rpa_policy
ON role_policy_assignments(policy_id);

CREATE TABLE user_policy_assignments (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL,
    resource_id BIGINT NOT NULL,
    policy_id BIGINT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_user_permission_policy
        UNIQUE (user_id, resource_id, policy_id),

    CONSTRAINT fk_upa_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_upa_permission
        FOREIGN KEY (resource_id)
        REFERENCES resources(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_upa_policy
        FOREIGN KEY (policy_id)
        REFERENCES policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_upa_lookup
ON user_policy_assignments(user_id, resource_id);

CREATE INDEX idx_upa_policy
ON user_policy_assignments(policy_id);

-- +goose Down

DROP TABLE IF EXISTS user_policy_assignments;
DROP TABLE IF EXISTS role_policy_assignments;
DROP TABLE IF EXISTS policies;

DROP TYPE IF EXISTS policy_effect;