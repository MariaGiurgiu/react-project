package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	_ = r.PathPrefix("/users").Subrouter()
	//userSubrouter.HandleFunc("/{id:[0-9]+}", handlers.UsersGetSingle).Methods(http.MethodGet)



	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return r
}