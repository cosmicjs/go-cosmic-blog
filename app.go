package main

import (
	"fmt"
	"go-cosmic-blog/src/handlers"
	"go-cosmic-blog/src/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()

	// Choose the folder to serve
	staticDir := "/static/"

	// Create the route
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.HandleFunc("/{slug}", handlers.Single).Methods("GET")

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	port := utils.GetPortEnv()

	fmt.Println("Starting server at port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
