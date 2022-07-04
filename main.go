package main

import (
	"fmt"
	//"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"credit-score-system/controllers"
	"os"
)

func handleRequests() {
	
	myRouter := mux.NewRouter().StrictSlash(true)
	
	
	myRouter.HandleFunc("/calculateCreditScore", controllers.Score).Methods("GET")
	
	//handle cors acces
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    })
	handler := c.Handler(myRouter)
	//port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":8900", handler))
}

func main() {
	fmt.Print("Credit Score System bit-v1")
	fmt.Println("Running on:: ",os.Getenv("PORT"))
	handleRequests()
}