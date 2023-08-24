package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Users struct {
	ID string
	Name string
	Age	uint8
}

var users []Users

func main() {
	r := mux.NewRouter()

	users = append(users, 
		Users{ID: "1", Name: "Aravind", Age: 20},
	)

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}