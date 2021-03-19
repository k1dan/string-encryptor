package main

import (
	"github.com/gorilla/mux"
	"github.com/k1dan/string-encryptor/randomizer/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.CreateStrings).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}

