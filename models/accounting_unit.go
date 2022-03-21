package models

import (
	"github.com/jinzhu/gorm"
)
// accounting_unit struct
type AccountingUnit struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
	TreasuryId int `json:"treasury_id" gorm:"index" binding:"required"`
	EntityId int `json:"entity_id" gorm:"index" binding:"required"`
	gorm.Model
	Entity Entity `gorm:"foreignkey:EntityId"`
	Treasury Treasury `gorm:"foreignkey:TreasuryId"`
}

type AccountingUnitIndexResponse struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
	gorm.Model
}

// type AccountingUnitWithRelationsResponse struct {
// 	NameEn string `json:"name_en" binding:"required" gorm:"unique"`
// 	Status string `json:"status" binding:"required" gorm:"default:'active'"`
// 	TreasuryId uint `json:"treasury_id" gorm:"index" binding:"required"`
// 	Treasury  Treasury
// 	EntityId uint `json:"entity_id" gorm:"index" binding:"required"`
// 	Entity  Entity
// 	gorm.Model
// }

// create accounting_unit input struct
type CreateAccountingUnitInput struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status"`
	// belongs to treasury
	TreasuryId int `json:"treasury_id" binding:"required" gorm:"foreignkey:TreasuryId"`
	// belongs to entity
	EntityId int `json:"entity_id" binding:"required"`
}

// update accounting_unit input struct
type UpdateAccountingUnitInput struct {
	NameEn string `json:"name_en"`
	Status string `json:"status"`
	// belongs to treasury
	TreasuryId int `json:"treasury_id"`
	// belongs to entity
	EntityId int `json:"entity_id"`
}