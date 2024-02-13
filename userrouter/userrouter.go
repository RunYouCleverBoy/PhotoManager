package userrouter

import (
	"encoding/json"
	"net/http"
	"time"
)

func QueryAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	users := []UserRequest{}
	json.NewEncoder(w).Encode(&users)
}

func QueryOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	user := UserRequest{
		ID:       id,
		Username: "user" + id,
		Name:     "User " + id,
		Token:    "1234567890" + time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(&user)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	userReq := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	time.Sleep(2 * time.Second) // Simulate a long process
	user := UserRequest{
		ID:       "User ID",
		Username: userReq.Username,
		Name:     userReq.Name,
		Token:    "1234567890" + time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(&user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	requestUser := UserRequest{}
	id := r.URL.Query().Get("id")
	json.NewDecoder(r.Body).Decode(&requestUser)
	go func() {
		response := UserRequest{
			ID:       id,
			Username: requestUser.Username,
			Name:     requestUser.Name,
			Token:    "1234567890" + time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(&response)
	}()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a user"))
	w.Write([]byte("User ID: " + r.URL.Query().Get("id")))
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserUpdateRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
