package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/arikarim/go-cfa/models"
	"github.com/arikarim/go-cfa/errors"
)

// get all accounting_units
func GetAccountingUnits(c *gin.Context) {
	var accounting_units []models.AccountingUnit
	if err := models.DB.Find(&accounting_units).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, accounting_units)
}

// create a new accounting_unit
func CreateAccountingUnit(c *gin.Context) {
	var input models.CreateAccountingUnitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// create a new accounting_unit
	accounting_unit := models.AccountingUnit{
		NameEn: input.NameEn,
		Status: input.Status,
		TreasuryId: input.TreasuryId,
		EntityId: input.EntityId,
	}

	// check if relations exists
	relations := map[uint]interface{}{
		accounting_unit.TreasuryId: &models.Treasury{},
		accounting_unit.EntityId: &models.Entity{},
	}
	if structName := errors.RecordNotFound(relations); structName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "relation " + structName + " not found",
		})
		return
	}

	// Start a transaction
	tx := models.DB.Begin()
	// create a new accounting_unit
	if err := tx.Create(&accounting_unit).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// commit transaction
	tx.Commit()

	c.JSON(200, gin.H{
		"message": "success",
		"data":    accounting_unit,
	})
}

// update a accounting_unit
func UpdateAccountingUnit(c *gin.Context) {
	var accounting_unit = models.AccountingUnit{}
	// find accounting_unit by id
	if err := models.DB.Where("id = ?", c.Param("id")).First(&accounting_unit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	// validate input
	var input models.UpdateAccountingUnitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// update a accounting_unit
	models.DB.Model(&accounting_unit).Updates(models.AccountingUnit{
		NameEn: input.NameEn,
		Status: input.Status,
		TreasuryId: input.TreasuryId,
		EntityId: input.EntityId,
	})

	c.JSON(200, gin.H{
		"message": "success",
		"data":    accounting_unit,
	})
}