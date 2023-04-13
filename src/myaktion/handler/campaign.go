package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nidigeser/go-myaktion/src/myaktion/model"
	"github.com/nidigeser/go-myaktion/src/myaktion/service"
	log "github.com/sirupsen/logrus"
)

func getCampaign(r *http.Request) (*model.Campaign, error) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Errorf("Can't decode request body to campaign struct: %v", err)
		return nil, err
	}
	return &campaign, nil
}

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign *model.Campaign
	campaign, err := getCampaign(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(campaign); err != nil {
		log.Printf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaign)
}

func GetCampaigns(w http.ResponseWriter, r *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Printf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaigns)
}
