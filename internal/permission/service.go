package permission

import (
	"context"	
)

type Service struct {
    queries *db.Queries
}

func NewService(q *db.Queries) *Service {
    return &Service{
        queries: q,
    }
}

func (s *Service) HasPermission(
    ctx context.Context,
    userID int64,
    resource string,
    action string,
) (bool, error) {

    return s.queries.HasPermission(ctx, db.HasPermissionParams{
        UserID:   userID,
        Resource: resource,
        Action:   action,
    })
}