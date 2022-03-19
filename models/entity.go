package models

// entity struct
type Entity struct {
	Id     uint `json:"id" gorm:"primary_key"`
	NameEn string `json:"name_en" binding:"required"`
	Status string `json:"status" binding:"required" gorm:"default:'active'"`
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
	