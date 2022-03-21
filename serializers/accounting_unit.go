package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/arikarim/go-cfa/models"
)

type AccountingUnitSerializer struct {
	C *gin.Context
	models.AccountingUnit
}

type AccountingUnitsSerializer struct {
	C *gin.Context
	AccountingUnits []models.AccountingUnit
}

type AccountingUnitIndexResponse struct {
	ID     uint
	NameEn string
	Status string
	CreatedAt string
	UpdatedAt string
}

type AccountingUnitProfileResponse struct {
	ID     uint
	NameEn string
	Status string
	CreatedAt string
	UpdatedAt string
	Entity EntityProfileResponse
	Treasury TreasuryProfileResponse
}

func (s *AccountingUnitSerializer) Response() AccountingUnitProfileResponse {
	entitySerializer := EntitySerializer{C: s.C, Entity: s.Entity}
	treasurySerializer := TreasurySerializer{C: s.C, Treasury: s.Treasury}
	response := AccountingUnitProfileResponse{
		ID:     s.AccountingUnit.ID,
		NameEn: s.AccountingUnit.NameEn,
		Status: s.AccountingUnit.Status,
		CreatedAt: s.AccountingUnit.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: s.AccountingUnit.UpdatedAt.Format("2006-01-02 15:04:05"),
		Entity: entitySerializer.Response(),
		Treasury: treasurySerializer.Response(),
	}

	return response
}

func (s *AccountingUnitsSerializer) Response() []AccountingUnitIndexResponse {
	var response []AccountingUnitIndexResponse
	for _, accountingUnit := range s.AccountingUnits {
		response = append(response, AccountingUnitIndexResponse{
			ID:     accountingUnit.ID,
			NameEn: accountingUnit.NameEn,
			Status: accountingUnit.Status,
			CreatedAt: accountingUnit.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: accountingUnit.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response
}
