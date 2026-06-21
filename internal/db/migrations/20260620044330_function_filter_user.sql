-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION scope_users(
    p_scope text,
    p_manager_id bigint DEFAULT NULL
)
RETURNS TABLE(user_id bigint)
LANGUAGE plpgsql
STABLE
AS $$
BEGIN

    --------------------------------------------------------------------
    -- SELF
    --------------------------------------------------------------------
    IF p_scope = 'SELF' THEN

        RETURN QUERY
        SELECT u.id
        FROM users u
        WHERE u.id = p_manager_id;

        RETURN;

    END IF;

    --------------------------------------------------------------------
    -- DIRECT REPORTS
    --------------------------------------------------------------------
    IF p_scope = 'DIRECT_REPORTS' THEN

        RETURN QUERY
        SELECT u.id
        FROM users u
        WHERE u.reporting_to = p_manager_id;

        RETURN;

    END IF;

    --------------------------------------------------------------------
    -- ALL USERS
    --------------------------------------------------------------------
    IF p_scope = 'ALL' THEN

        RETURN QUERY
        SELECT u.id
        FROM users u;

        RETURN;

    END IF;

    --------------------------------------------------------------------
    -- ALL REPORTS (RECURSIVE)
    --------------------------------------------------------------------
    RETURN QUERY

    WITH RECURSIVE reports(user_id) AS (

        --------------------------------------------------------
        -- Direct reports
        --------------------------------------------------------
        SELECT u.id
        FROM users u
        WHERE u.reporting_to = p_manager_id

        UNION ALL

        --------------------------------------------------------
        -- Child reports
        --------------------------------------------------------
        SELECT child.id
        FROM users child
        JOIN reports parent
            ON child.reporting_to = parent.user_id
    )

    SELECT DISTINCT reports.user_id
    FROM reports;

END;
$$;


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

SELECT DISTINCT u.id
FROM scope_users(
        p_scope,
        p_manager_id
     ) s

JOIN users u
    ON u.id = s.user_id

JOIN roles r
    ON r.id = u.role_id

WHERE

    --------------------------------------------------------------------
    -- OFFICE FILTER
    --------------------------------------------------------------------
    (
        p_office_ids IS NULL
        OR cardinality(p_office_ids) = 0
        OR u.office_id = ANY(p_office_ids)
    )

AND

    --------------------------------------------------------------------
    -- DEPARTMENT FILTER
    --------------------------------------------------------------------
    (
        p_department_ids IS NULL
        OR cardinality(p_department_ids) = 0
        OR u.department_id = ANY(p_department_ids)
    )

AND

    --------------------------------------------------------------------
    -- HIERARCHY FILTER
    --------------------------------------------------------------------
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

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin

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

DROP FUNCTION IF EXISTS scope_users(
    text,
    bigint
);

-- +goose StatementEnd
