package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetGuitar(w http.ResponseWriter, r *http.Request) {
	weapons := LoadGuitarsCsv()
	json.NewEncoder(w).Encode(weapons)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/guitars", GetGuitar).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
