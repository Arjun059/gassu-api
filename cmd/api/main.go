package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"context"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

import "gassu/internal/db/sqlc"
import "github.com/jackc/pgx/v5/pgtype"

func main() {
	cwd, _ := os.Getwd()

	err := godotenv.Load(path.Join(cwd, "./internal/config", "local.env"))
	if err != nil {
		log.Fatalf("Failed to load evn: %v", err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx,
		"postgres://postgres:arjun@localhost:5433/test")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := sqlc.New(pool)

	r := mux.NewRouter()

	r.HandleFunc("/user/create",func(w http.ResponseWriter, r *http.Request) {
	user, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
    Name: pgtype.Text{
        String: "Arjun",
        Valid:  true,
    },
    Email: pgtype.Text{
        String: "arjun@example.com",
        Valid:  true,
    },
})

if err != nil {
    log.Printf("Create error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

fmt.Printf("%+v\n", user)
	}).Methods("GET")

	r.HandleFunc("/user/get",func(w http.ResponseWriter, r *http.Request) {
		user, err := queries.GetUser(ctx, 1)

		if err != nil {
			log.Printf("GetUser error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v\n", user)
		fmt.Fprintf(w, "Get user route %+v", user)
	}).Methods("GET")

	r.HandleFunc("/protected",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello Protected Route")
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux!"))
	})

	fmt.Println("Server running at: http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
