-- +goose Up

CREATE OR REPLACE FUNCTION filter_users(
    p_scope text,
    p_manager_id bigint DEFAULT NULL,

    p_office_ids bigint[] DEFAULT NULL,
    p_department_ids bigint[] DEFAULT NULL,
    p_company_ids bigint[] DEFAULT NULL,
    p_employment_types text[] DEFAULT NULL,

    p_my_hierarchy bigint DEFAULT NULL,
    p_hierarchy_mode text DEFAULT NULL
)
RETURNS TABLE(user_id bigint)
LANGUAGE sql
STABLE
AS $$

WITH RECURSIVE scope_users AS (

    ----------------------------------------------------
    -- FAST PATH: ALL USERS (NO RECURSION)
    ----------------------------------------------------
    SELECT u.id
    FROM users u
    WHERE p_scope = 'ALL'

    UNION ALL

    ----------------------------------------------------
    -- SELF
    ----------------------------------------------------
    SELECT u.id
    FROM users u
    WHERE p_scope = 'SELF'
      AND u.id = p_manager_id

    UNION ALL

    ----------------------------------------------------
    -- DIRECT REPORTS
    ----------------------------------------------------
    SELECT u.id
    FROM users u
    WHERE p_scope = 'DIRECT_REPORTS'
      AND u.reporting_manager_id = p_manager_id

    UNION ALL

    ----------------------------------------------------
    -- ALL_REPORTS START
    ----------------------------------------------------
    SELECT u.id
    FROM users u
    WHERE p_scope = 'ALL_REPORTS'
      AND u.reporting_manager_id = p_manager_id

    UNION ALL

    ----------------------------------------------------
    -- RECURSION ONLY FOR ALL_REPORTS
    ----------------------------------------------------
    SELECT child.id
    FROM users child
    JOIN scope_users parent
        ON child.reporting_manager_id = parent.id
    WHERE p_scope = 'ALL_REPORTS'
)

SELECT DISTINCT u.id
FROM users u
JOIN scope_users s ON s.id = u.id
JOIN roles r ON r.id = u.role_id

WHERE

    ----------------------------------------------------
    -- OFFICE FILTER
    ----------------------------------------------------
    (
        p_office_ids IS NULL
        OR cardinality(p_office_ids) = 0
        OR u.office_id = ANY(p_office_ids)
    )

AND

    ----------------------------------------------------
    -- DEPARTMENT FILTER
    ----------------------------------------------------
    (
        p_department_ids IS NULL
        OR cardinality(p_department_ids) = 0
        OR u.department_id = ANY(p_department_ids)
    )

AND

    ----------------------------------------------------
    -- COMPANY FILTER
    ----------------------------------------------------
    (
        p_company_ids IS NULL
        OR cardinality(p_company_ids) = 0
        OR u.company_id = ANY(p_company_ids)
    )

AND

    ----------------------------------------------------
    -- EMPLOYMENT TYPE FILTER
    ----------------------------------------------------
    (
        p_employment_types IS NULL
        OR cardinality(p_employment_types) = 0
        OR u.employment_type = ANY(p_employment_types)
    )

AND

    ----------------------------------------------------
    -- HIERARCHY RULE
    ----------------------------------------------------
    (
        p_my_hierarchy IS NULL
        OR p_hierarchy_mode IS NULL
        OR (
            p_hierarchy_mode = 'LOWER_ONLY'
            AND r.hierarchy < p_my_hierarchy
        )
        OR (
            p_hierarchy_mode = 'SAME_AND_LOWER'
            AND r.hierarchy <= p_my_hierarchy
        )
    );

$$;


-- +goose Down

DROP FUNCTION IF EXISTS filter_users(
    text,
    bigint,
    bigint[],
    bigint[],
    bigint[],
    text[],
    bigint,
    text
);