
-- +goose Up

ALTER TABLE roles
DROP CONSTRAINT roles_hierarchy_key;


-- +goose Down

ALTER TABLE roles
ADD CONSTRAINT roles_hierarchy_key UNIQUE (hierarchy);

