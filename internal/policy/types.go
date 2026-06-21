package policy

type Scope string

const (
	ScopeSelf          Scope = "SELF"
	ScopeDirectReports Scope = "DIRECT_REPORTS"
	ScopeAllReports    Scope = "ALL_REPORTS"
	ScopeAll           Scope = "ALL"
)

type HierarchyMode string

const (
	HierarchyLowerOnly    HierarchyMode = "LOWER_ONLY"
	HierarchySameAndLower HierarchyMode = "SAME_AND_LOWER"
)

// Policy is loaded from database.
type Policy struct {
	Scope Scope

	OfficeIDs       []int64
	DepartmentIDs   []int64
	CompanyIDs      []int64
	EmploymentTypes []string

	HierarchyMode *HierarchyMode
}

// Filter is the final object that will be passed to sqlc.
type Filter struct {
	Scope Scope

	ManagerID *int64

	OfficeIDs       []int64
	DepartmentIDs   []int64
	CompanyIDs      []int64
	EmploymentTypes []string

	MyHierarchy   *int64
	HierarchyMode *HierarchyMode
}
