package users

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

type UserHandler struct {
	Querier sqlc.Querier
}

func NewHandler(querier sqlc.Querier) *UserHandler {
	return &UserHandler{
		Querier: querier,
	}
}

func (h *UserHandler) createUser(c echo.Context) error {
	ctx := context.Background()

	reqUser := new(ReqUser)
	if err := c.Bind(reqUser); err != nil {
		return err
	}

	user, err := h.Querier.CreateUser(ctx, sqlc.CreateUserParams{
		Name:   reqUser.Name,
		RoleID: reqUser.RoleID,
	})

	if err != nil {
		log.Printf("Create error: %v", err)
		return err
	}

	fmt.Printf("%+v\n", user)
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) getUser(c echo.Context) error {
	ctx := context.Background()

	userID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := h.Querier.GetUser(ctx, userID)

	allowed, perErro := h.Querier.HasPermission(ctx, sqlc.HasPermissionParams{
		UserID:   32,
		Resource: "employee",
		Action:   "read",
	})

	fmt.Println("allowed: ", allowed)
	fmt.Println("perErro: ", perErro)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No user found!",
			})
		}
		log.Printf("get error: %v", err)
		return err
	}

	fmt.Printf("%+v\n", user)
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) updateUser(c echo.Context) error {
	ctx := context.Background()

	reqUser := new(ReqUser)
	if err := c.Bind(reqUser); err != nil {
		return c.String(http.StatusBadRequest, "Req body invalid!")
	}

	h.Querier.UpdateUser(ctx, sqlc.UpdateUserParams{
		Name:   reqUser.Name,
		RoleID: reqUser.RoleID,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"error":   false,
		"message": "User updated successfully!",
	})
}

func (h *UserHandler) getUserList(c echo.Context) error {
	ctx := context.Background()

	userID, _ := strconv.ParseInt(c.QueryParam("id"), 10, 64)

	allowed, perErro := h.Querier.HasPermission(ctx, sqlc.HasPermissionParams{
		UserID:   userID,
		Resource: "employee",
		Action:   "read",
	})

	fmt.Println("allowed: ", allowed)
	fmt.Println("perErro: ", perErro)

	users, err := h.Querier.ListUsers(ctx)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No user found!",
			})
		}
		log.Printf("get error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, users)
}
