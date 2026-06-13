package roles

import (
	"gassu/internal/db/sqlc"

	"github.com/labstack/echo"
)

type RoleModule struct {
	Querier sqlc.Querier
	Echo    *echo.Echo
}

func NewRoleModule(echo *echo.Echo, querier sqlc.Querier) *RoleModule {
	return &RoleModule{
		Querier: querier,
		Echo:    echo,
	}
}

func (m *RoleModule) RegisterRoutes() {

	roleHandler := NewRoleHandler(m.Querier)
	g := m.Echo.Group("/roles")

	g.POST("", roleHandler.createRole)
	g.GET("/:id", roleHandler.getRole)

}
