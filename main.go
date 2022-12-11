package main

import (
	"blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
