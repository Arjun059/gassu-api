-- +goose Up

CREATE TABLE policy_rules (
    id BIGSERIAL PRIMARY KEY,

    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,

    parent_rule_id BIGINT REFERENCES policy_rules(id),

    logical_operator VARCHAR(3) NOT NULL CHECK (
        logical_operator IN ('AND', 'OR')
    )
);

CREATE TABLE policy_conditions (
    id BIGSERIAL PRIMARY KEY,

    rule_id BIGINT NOT NULL REFERENCES policy_rules(id),

    field_name VARCHAR(100) NOT NULL,

    operator VARCHAR(20) NOT NULL,

    value_type VARCHAR(30) NOT NULL,

    value TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS policy_rules;
DROP TABLE IF EXISTS policy_conditions;
