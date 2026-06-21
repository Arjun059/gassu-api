package policy

import "context"

type Repository interface {
	GetPolicy(
		ctx context.Context,
		userID int64,
		resource string,
		action string,
	) (*Policy, error)

	GetUserHierarchy(
		ctx context.Context,
		userID int64,
	) (*int64, error)
}
