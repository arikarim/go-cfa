package models

import (
	"github.com/jinzhu/gorm"
)

// entity struct
type Entity struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
	gorm.Model
	AccountingUnits []AccountingUnit
}

// create entity input struct
type CreateEntityInput struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status"`
}

// update entity input struct
type UpdateEntityInput struct {
	NameEn string `json:"name_en"`
	Status string `json:"status"`
}
	