package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model

	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Type        string `gorm:"not null"`
	ProofImage  string `json:"proof_image"`
	Status      string `gorm:"default:Pending"`
	Credits     int    `gorm:"default:0"`
}
