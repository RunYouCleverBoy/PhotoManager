package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"playgrounds.com/userrouter"
)

func Listen() {
	r := mux.NewRouter()
	var api *mux.Router = r.PathPrefix("/api").Subrouter()
	var user = api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/queryall", userrouter.QueryAll)
	user.HandleFunc("/queryone/{id}", userrouter.QueryOne)
	user.HandleFunc("/create", userrouter.Create).Methods("POST")
	user.HandleFunc("/update", userrouter.Update).Methods("POST")
	user.HandleFunc("/delete", userrouter.Delete).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
