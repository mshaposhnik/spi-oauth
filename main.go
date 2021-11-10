package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mshaposhnik/spi-oauth/controllers"
	"os"
	"fmt"
)


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/authenticate", controllers.Authenticate).Methods("GET")
	router.HandleFunc("/callback", controllers.Callback).Methods("GET")

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
