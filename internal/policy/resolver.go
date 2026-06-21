package policy

import (
	"context"
	"encoding/json"
	"fmt"
	"gassu/internal/db/sqlc"
)

type PolicyRules struct {
	Scope           string   `json:"scope"`
	OfficeIds       []int64  `json:"office_ids"`
	DepartmentIds   []int64  `json:"department_ids"`
	CompanyIds      []int64  `json:"company_ids"`
	EmploymentTypes []string `json:"employment_types"`
	HierarchyMode   string   `json:"hierarchy"`
}

func Resolve(
	ctx context.Context,
	q sqlc.Querier,
	userID int64,
	resourceID int64,
) (*Filter, error) {

	hierarchy, err := q.GetUserHierarchy(ctx, userID)
	if err != nil {
		return nil, err
	}

	policies, err := q.GetPolicies(ctx, sqlc.GetPoliciesParams{
		UserID:     userID,
		ResourceID: resourceID,
	})
	if err != nil {
		return nil, err
	}

	if len(policies) == 0 {
		return nil, nil // or return default filter
	}

	policy := policies[0]

	var rules PolicyRules
	marshelErro := json.Unmarshal(policy.Rules, &rules)

	fmt.Println(marshelErro)

	return &Filter{
		Scope: Scope(rules.Scope),

		ManagerID: &userID,

		OfficeIDs:       rules.OfficeIds,
		DepartmentIDs:   rules.DepartmentIds,
		CompanyIDs:      rules.CompanyIds,
		EmploymentTypes: rules.EmploymentTypes,

		MyHierarchy: &hierarchy,
	}, nil
}
