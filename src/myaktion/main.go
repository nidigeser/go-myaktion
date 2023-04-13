package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nidigeser/go-myaktion/src/myaktion/handler"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/helath", handler.Health).Methods("GET")
	router.HandleFunc("/campaign", handler.CreateCampaign).Methods("POST")
	router.HandleFunc("/campaigns", handler.GetCampaigns).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
