package models

// type Status int

import (
	"github.com/jinzhu/gorm"
)

// Treasury struct
type Treasury struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
	gorm.Model
	AccountingUnits []AccountingUnit
}

// CreateTreasuryInput struct
type CreateTreasuryInput struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status"`
}

type UpdateTreasuryInput struct {
	NameEn string `json:"name_en"`
	Status string `json:"status"`
}


