package users

import (
	"context"
	"errors"
	"fmt"
	"gassu/internal/auth"
	"gassu/internal/db/sqlc"
	"gassu/internal/policy"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
		Name: reqUser.Name,
		RoleID: pgtype.Int8{
			Int64: reqUser.RoleID,
			Valid: true,
		},
	})

	if err != nil {
		log.Printf("Create error: %v", err)
		return err
	}

	token, _ := auth.GenerateToken(user.ID)

	fmt.Printf("%+v\n", CreateUserResponseDTO{Token: token, UserID: user.ID})
	return c.JSON(http.StatusOK, CreateUserResponseDTO{Token: token, UserID: user.ID})
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
		Name: reqUser.Name,
		RoleID: pgtype.Int8{
			Int64: reqUser.RoleID,
			Valid: true,
		},
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

	perm, err := h.Querier.HasPermission(ctx, sqlc.HasPermissionParams{
		UserID:   userID,
		Resource: "users",
		Action:   "read",
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "permission denied")
	}

	if !perm.HasPermission {
		return echo.NewHTTPError(http.StatusForbidden, "permission denied")
	}

	fmt.Printf("Perm %+v", perm)
	// Access user data
	fmt.Println(perm.ID)
	fmt.Println(perm.Name)

	filter, err := policy.Resolve(
		ctx,
		h.Querier,
		userID,
		perm.ResourceID,
	)

	if err != nil {
		return err
	}

	users, err := h.Querier.ListUsers(ctx, sqlc.ListUsersParams{
		Scope:           string(filter.Scope),
		ManagerID:       pgtype.Int8{Int64: *filter.ManagerID, Valid: true},
		OfficeIds:       filter.OfficeIDs,
		DepartmentIds:   filter.DepartmentIDs,
		CompanyIds:      filter.CompanyIDs,
		EmploymentTypes: filter.EmploymentTypes,
		MyHierarchy:     pgtype.Int8{Int64: *filter.MyHierarchy, Valid: true},
		HierarchyMode:   pgtype.Text{String: string(*filter.HierarchyMode), Valid: filter.HierarchyMode != nil},
	})

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
