package permission

import (
	"context"
	"errors"
	"fmt"
	"gassu/internal/db/sqlc"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo"
)

type PermissionHandler struct {
	Querier sqlc.Querier
}

func NewHandler(querier sqlc.Querier) *PermissionHandler {
	return &PermissionHandler{
		Querier: querier,
	}
}

func (h *PermissionHandler) createPermission(c echo.Context) error {
	ctx := context.Background()

	reqPerm := new(ReqPermission)
	if err := c.Bind(reqPerm); err != nil {
		return err
	}

	user, err := h.Querier.CreatePermission(ctx, sqlc.CreatePermissionParams{
		Resource: reqPerm.Resource,
		Action:   reqPerm.Action,
	})

	if err != nil {
		log.Printf("Create error: %v", err)
		return err
	}

	fmt.Printf("%+v\n", user)
	return c.JSON(http.StatusOK, user)
}

func (h *PermissionHandler) getPermission(c echo.Context) error {
	ctx := context.Background()

	perId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	fmt.Println("permission perId:", perId)

	perm, err := h.Querier.GetPermission(ctx, perId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Perm not found!",
			})
		}
		log.Printf("get error: %v", err)
		return err
	}

	fmt.Printf("%+v\n", perm)
	return c.JSON(http.StatusOK, perm)
}

func (h *PermissionHandler) updatePermission(c echo.Context) error {
	ctx := context.Background()

	reqPerm := new(ReqPermission)
	if err := c.Bind(reqPerm); err != nil {
		return c.String(http.StatusBadRequest, "Req body invalid!")
	}

	h.Querier.UpdatePermission(ctx, sqlc.UpdatePermissionParams{
		Resource: reqPerm.Resource,
		Action:   reqPerm.Action,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"error":   false,
		"message": "User updated successfully!",
	})
}
