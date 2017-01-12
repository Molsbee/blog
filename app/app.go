package app

import (
	"net/http"
	"github.com/molsbee/blog/service"
	"github.com/gorilla/mux"
	"log"
)

func Start(configuration service.Configuration) {
	router := mux.NewRouter();
	log.Println(http.ListenAndServe(":"+configuration.Port, router))
}
