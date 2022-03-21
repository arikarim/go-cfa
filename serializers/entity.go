package serializers

import (
	"github.com/arikarim/go-cfa/models"
	"github.com/gin-gonic/gin"
)

type EntitySerializer struct {
	C *gin.Context
	models.Entity
}

type EntitiesSerializer struct {
	C *gin.Context
	Entities []models.Entity
}

type EntityIndexResponse struct {
	ID     uint
	NameEn string
	Status string
	CreatedAt string
	UpdatedAt string
}

type EntityProfileResponse struct {
	ID     uint
	NameEn string
	Status string
	CreatedAt string
	UpdatedAt string
	AccountingUnits []AccountingUnitIndexResponse
}

func (s *EntitySerializer) Response() EntityProfileResponse {
	accountingUnitSerializer := AccountingUnitsSerializer{
		C: s.C,
		AccountingUnits: s.AccountingUnits,
	}
	response := EntityProfileResponse{
		ID:     s.Entity.ID,
		NameEn: s.Entity.NameEn,
		Status: s.Entity.Status,
		CreatedAt: s.Entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: s.Entity.UpdatedAt.Format("2006-01-02 15:04:05"),
		AccountingUnits: accountingUnitSerializer.Response(),
	}

	return response
}

func (s *EntitiesSerializer) Response() []EntityIndexResponse {
	var response []EntityIndexResponse
	for _, entity := range s.Entities {
		response = append(response, EntityIndexResponse{
			ID:     entity.ID,
			NameEn: entity.NameEn,
			Status: entity.Status,
			CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response
}
