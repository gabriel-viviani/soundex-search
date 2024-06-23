package model

import "gorm.io/gorm"

type SanctionEntity struct {
	gorm.Model
	LogicalID int `gorm:"index"`
	Alias     string
}

func (SanctionEntity) TableName() string {
	return "sanction_entities"
}
