package main

import (
	"log"
	"net/http"
	"time"

	"go-simple-clean-architecture/api/router"
)

func main () {

	srv := &http.Server{
        Handler:      router.ConfigureRoutes(),
        Addr:         "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	log.Fatal(srv.ListenAndServe())
}