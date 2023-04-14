package service

import (
	"github.com/nidigeser/go-myaktion/src/myaktion/db"
	"github.com/nidigeser/go-myaktion/src/myaktion/model"
	log "github.com/sirupsen/logrus"
)

func AddDonation(campaignId uint, donation *model.Donation) error {
	_, err := GetCampaign(campaignId)
	if err != nil {
		return err
	}
	donation.CampaignID = campaignId
	result := db.DB.Create(donation)
	if result.Error != nil {
		return result.Error
	}
	entry := log.WithField("ID", campaignId)
	entry.Info("Successfully added new donation to campaign in database.")
	entry.Tracef("Stored: %v", donation)
	return nil
}
