package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range users {
		if params["id"] == user.ID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var user Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, existingUser := range users {
		if user.ID == existingUser.ID {
			http.Error(w, "Duplicate ID", http.StatusBadRequest)
			return
		}
	}

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index + 1:]...)
			var user Users
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index + 1:]...)
			json.NewEncoder(w).Encode("Deleted User")
			break
		}
	}
}