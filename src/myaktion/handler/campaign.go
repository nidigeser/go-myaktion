package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nidigeser/go-myaktion/src/myaktion/model"
	"github.com/nidigeser/go-myaktion/src/myaktion/service"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Printf("Cannot decode request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(&campaign); err != nil {
		log.Printf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaign); err != nil {
		log.Printf("Failure encoding campaign to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCampaigns(w http.ResponseWriter, r *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Printf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaigns); err != nil {
		log.Printf("Failure encoding campaigns to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
