package users

import (
	"gassu/internal/db/sqlc"

	"github.com/labstack/echo"
)

type UserModule struct {
	echo    *echo.Echo
	querier sqlc.Querier
}

func NewUserModule(echo *echo.Echo, querier sqlc.Querier) *UserModule {
	return &UserModule{
		echo:    echo,
		querier: querier,
	}
}

func (m *UserModule) RegisterRoutes() {

	handler := NewHandler(m.querier)

	g := m.echo.Group("/users")
	g.POST("", handler.createUser)
	g.GET("/:id", handler.getUser)
}
