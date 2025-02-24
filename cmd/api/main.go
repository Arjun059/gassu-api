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

	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Auto-migrate  schema's
	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
		&models.Product{},
	)

	userHandler := &handlers.UserHandler{DB: db}
	blogHandler := &handlers.BlogHandler{DB: db}
	productHandler := &handlers.ProductHandler{DB: db}

	r := mux.NewRouter()

	r.HandleFunc("/user/get/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/user/create", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/delete/{id}", userHandler.DeleteUser).Methods("DELETE")

	r.HandleFunc("/user/sign-up", userHandler.SignupUser).Methods("POST")
	r.HandleFunc("/user/sign-in", userHandler.SigninUser).Methods("POST")

	r.HandleFunc("/blog/create", utils.WithAuth(blogHandler.CreateBlog)).Methods("POST")
	r.HandleFunc("/blog/get/{id}", utils.WithAuth(blogHandler.GetBlog)).Methods("GET")
	r.HandleFunc("/blog/update/{id}", utils.WithAuth(blogHandler.UpdateBlog)).Methods("PUT")
	r.HandleFunc("/blog/delete/{id}", utils.WithAuth(blogHandler.DeleteBlog)).Methods("DELETE")
	r.HandleFunc("/blog/list", utils.WithAuth(blogHandler.ListBlogs)).Methods("GET")

	r.HandleFunc("/product/create", utils.WithAuth(productHandler.CreateProduct)).Methods("POST")
	r.HandleFunc("/product/update/{id}", utils.WithAuth(productHandler.UpdateProduct)).Methods("PUT")
	r.HandleFunc("/product/get/{id}", utils.WithAuth(productHandler.GetProduct)).Methods("GET")
	r.HandleFunc("/product/delete/{id}", utils.WithAuth(productHandler.DeleteProduct)).Methods("DELETE")
	r.HandleFunc("/product/list", utils.WithAuth(productHandler.ListProduct)).Methods("GET")

	r.HandleFunc("/protected", utils.WithAuth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello Protected Route")
	})).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gorilla Mux!"))
	})

	fmt.Println("Server running at: http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
