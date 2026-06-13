package permission

import (
	"gassu/internal/db/sqlc"

	"github.com/labstack/echo"
)

type PermissionModule struct {
	echo    *echo.Echo
	querier sqlc.Querier
}

func NewPermissionModule(echo *echo.Echo, querier sqlc.Querier) *PermissionModule {
	return &PermissionModule{
		echo:    echo,
		querier: querier,
	}
}

func (m *PermissionModule) RegisterRoutes() {
	handler := NewHandler(m.querier)

	g := m.echo.Group("/permissions")
	g.POST("", handler.createPermission)
	g.GET("/:id", handler.getPermission)
}
