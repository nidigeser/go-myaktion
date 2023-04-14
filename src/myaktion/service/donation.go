package service

import (
	"github.com/nidigeser/go-myaktion/src/myaktion/model"
	log "github.com/sirupsen/logrus"
)

func AddDonation(campaignId uint, donation *model.Donation) error {
	campaign, err := GetCampaign(campaignId)
	if err != nil {
		return err
	}
	campaign.Donations = append(campaign.Donations, *donation)
	entry := log.WithField("ID", campaignId)
	entry.Info("Successfully added new donation to campaign.")
	entry.Tracef("Stored: %v", donation)
	return nil
}
