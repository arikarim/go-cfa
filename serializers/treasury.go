package serializers

import (
	"github.com/arikarim/go-cfa/models"
	"github.com/gin-gonic/gin"
)

type TreasurySerializer struct {
	C *gin.Context
	models.Treasury
}

type TreasuriesSerializer struct {
	C *gin.Context
	Treasuries []models.Treasury
}

type TreasuryIndexResponse struct {
	ID     uint
	NameEn string
	Status string
	CreatedAt string
	UpdatedAt string
}

type TreasuryProfileResponse struct {
	ID     uint
	NameEn string
	Status string
	CreatedAt string
	UpdatedAt string
	AccountingUnits []AccountingUnitIndexResponse
}

func (s *TreasurySerializer) Response() TreasuryProfileResponse {
	accountingUnitSerializer := AccountingUnitsSerializer{
		C: s.C,
		AccountingUnits: s.AccountingUnits,
	}
	response := TreasuryProfileResponse{
		ID:     s.Treasury.ID,
		NameEn: s.Treasury.NameEn,
		Status: s.Treasury.Status,
		CreatedAt: s.Treasury.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: s.Treasury.UpdatedAt.Format("2006-01-02 15:04:05"),
		AccountingUnits: accountingUnitSerializer.Response(),
	}

	return response
}

func (s *TreasuriesSerializer) Response() []TreasuryIndexResponse {
	var response []TreasuryIndexResponse
	for _, Treasury := range s.Treasuries {
		response = append(response, TreasuryIndexResponse{
			ID:     Treasury.ID,
			NameEn: Treasury.NameEn,
			Status: Treasury.Status,
			CreatedAt: Treasury.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: Treasury.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response
}
