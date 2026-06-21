package main

import (
	"fmt"
	"log"
	"net/http"

	"context"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	"gassu/internal/config"
	"gassu/internal/db/sqlc"

	auth "gassu/internal/auth"
	"gassu/internal/domain/resources"
	"gassu/internal/domain/roles"
	"gassu/internal/domain/users"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// decide which env to use
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Failed to load evn: %v", err)
	}

	config, _ := config.Load()
	ctx := context.Background()

	pgxConfig, err := pgxpool.ParseConfig(config.POSTGRES_DB_URI)
	if err != nil {
		panic(err)
	}

	// Set the PostgreSQL session timezone to UTC
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := sqlc.New(pool)

	var tz string
	_ = pool.QueryRow(ctx, "SHOW TIME ZONE").Scan(&tz)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	api := e.Group("/api")
	api.Use(auth.JWTMiddleware)

	publicApi := e.Group("")

	users.NewUserModule(publicApi, queries).RegisterRoutes()
	roles.NewRoleModule(api, queries).RegisterRoutes()
	resources.NewPermissionModule(api, queries).RegisterRoutes()

	e.GET("/", func(c echo.Context) error {
		user, err := queries.GetUser(ctx, 1)

		if err != nil {
			log.Printf("GetUser error: %v", err)
			return err
		}

		fmt.Printf("%+v\n", user)
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println("Server running at: http://localhost:8000")
	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
