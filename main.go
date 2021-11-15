package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"spi-oauth/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/github/authenticate", controllers.GitHubAuthenticate).Methods("GET")
	router.HandleFunc("/github/callback", controllers.GitHubCallback).Methods("GET")

	router.HandleFunc("/quay/authenticate", controllers.QuayAuthenticate).Methods("GET")
	router.HandleFunc("/quay/callback", controllers.QuayCallback).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
