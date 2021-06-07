package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
