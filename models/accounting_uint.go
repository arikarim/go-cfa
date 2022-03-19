package models

// accounting_unit struct
type AccountingUnit struct {
	Id     uint `json:"id" gorm:"primary_key"`
	NameEn string `json:"name_en" binding:"required" gorm:"unique"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
	// belongs to treasury
	TreasuryId uint `json:"treasury_id" gorm:"index" binding:"required"`
	// Treasury  Treasury
	// belongs to entity
	EntityId uint `json:"entity_id" gorm:"index" binding:"required"`
	// Entity   Entity
}

// create accounting_unit input struct
type CreateAccountingUnitInput struct {
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status"`
	// belongs to treasury
	TreasuryId uint `json:"treasury_id" binding:"required" gorm:"foreignkey:TreasuryId"`
	// belongs to entity
	EntityId uint `json:"entity_id" binding:"required"`
}

// update accounting_unit input struct
type UpdateAccountingUnitInput struct {
	NameEn string `json:"name_en"`
	Status string `json:"status"`
	// belongs to treasury
	TreasuryId uint `json:"treasury_id"`
	// belongs to entity
	EntityId uint `json:"entity_id"`
}