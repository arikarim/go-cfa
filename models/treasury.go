package models

// type Status int

// // create status enum
// const (
// 	Active Status = iota
// 	Inactive
// )

// Treasury struct
type Treasury struct {
	Id     uint `json:"id" gorm:"primary_key"`
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
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


