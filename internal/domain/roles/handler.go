package roles

import (
	"context"
	"fmt"
	"gassu/internal/db/sqlc"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type RoleHandler struct {
	Querier sqlc.Querier
}

func NewRoleHandler(querier sqlc.Querier) *RoleHandler {
	return &RoleHandler{
		Querier: querier,
	}
}

func (h *RoleHandler) createRole(c echo.Context) error {
	ctx := context.Background()

	reqRole := new(ReqRole)

	if err := c.Bind(reqRole); err != nil {
		return err
	}

	role, err := h.Querier.CreateRole(ctx, sqlc.CreateRoleParams{
		Name:      reqRole.Name,
		Hierarchy: reqRole.Hierarchy,
	})

	if err != nil {
		fmt.Printf("error on create role %+v", err)
		return err
	}

	return c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) getRole(c echo.Context) error {
	ctx := context.Background()

	roleId, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)

	if parseErr != nil {
		fmt.Printf("parse error %+v", parseErr)
		return parseErr
	}

	role, err := h.Querier.GetRole(ctx, roleId)

	if err != nil {
		fmt.Printf("db error %+v", err)
		return err
	}

	return c.JSON(http.StatusOK, role)
}
