package model

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Name               string     `gorm:"notNull;size:30"`
	OrganizerName      string     `gorm:"notNull"`
	TargetAmount       float64    `gorm:"notNull;check:target_amount >= 10.0"`
	DonationMinimum    float64    `gorm:"notNull;check:donation_minimum >= 1.0"`
	Donations          []Donation `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,onDelete:CASCADE"`
	AmountDonatedSoFar float64    `gorm:"-"` //shall not be stored in Database
	Account            Account    `gorm:"embedded;embeddedPrefix:account_"`
}
