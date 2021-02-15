package router

import (
	"fmt"
	handler "go-simple-clean-architecture/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

//ConfigureRoutes configure routes
func ConfigureRoutes() (http.Handler){
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {fmt.Fprintf(w, "Running")}).Methods("GET", "POST")
    r.HandleFunc("/Users", handler.UserHandler).Methods("GET", "POST")
    r.HandleFunc("/Customers", handler.CustomerHandler).Methods("GET", "POST")
	
	return r
}