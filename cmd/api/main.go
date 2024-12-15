package main

import (
	"fmt"
	handlers "gassu/internal/handlers"
	"gassu/internal/models"
	"gassu/internal/utils"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	cwd, _ := os.Getwd()

	err := godotenv.Load(path.Join(cwd, "config", "local.env"))
	if err != nil {
		log.Fatalf("Failed to load evn: %v", err)
	}

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Auto-migrate  schema's
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Blog{})

	userHandler := &handlers.UserHandler{DB: db}
	blogHandler := &handlers.BlogHandler{DB: db}

	r := mux.NewRouter()

	r.HandleFunc("/user/get/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/user/create", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/delte/{id}", userHandler.DeleteUser).Methods("DELETE")

	r.HandleFunc("/user/signup", userHandler.SignupUser).Methods("POST")
	r.HandleFunc("/user/signin", userHandler.SigninUser).Methods("POST")

	r.HandleFunc("/blog/create", blogHandler.CreateBlog).Methods("POST")
	r.HandleFunc("/blog/update/{id}", blogHandler.UpdateBlog).Methods("PUT")
	r.HandleFunc("/blog/get/{id}", blogHandler.GetBlog).Methods("GET")
	r.HandleFunc("/blog/delete/{id}", blogHandler.DeleteBlog).Methods("DELETE")
	r.HandleFunc("/blog/list", blogHandler.ListBlogs).Methods("GET")

	r.HandleFunc("/protected", utils.WithAuth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello Prortected Route")
	})).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux!"))
	})
	http.ListenAndServe(":8080", r)
}
