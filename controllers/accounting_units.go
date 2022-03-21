package controllers

import (
	"net/http"

	"github.com/arikarim/go-cfa/errors"
	"github.com/arikarim/go-cfa/models"
	"github.com/arikarim/go-cfa/serializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	// Initialize seriaizer
	serializer := serializers.AccountingUnitsSerializer{
		C: c,
		AccountingUnits: accounting_units,
	}

	c.JSON(http.StatusOK, serializer.Response())
}

// get a accounting_unit by id
func GetAccountingUnit(c *gin.Context) {
	var accounting_unit = models.AccountingUnit{}
	// find accounting_unit by id
	if err := models.DB.Preload("Treasury").Preload("Entity").Where("id = ?", c.Param("id")).First(&accounting_unit).Error; err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Initialize seriaizer
	serializer := serializers.AccountingUnitSerializer{
		C: c,
		AccountingUnit: accounting_unit,
	}

	c.JSON(http.StatusOK, serializer.Response())
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
		NameEn:     input.NameEn,
		Status:     input.Status,
		TreasuryId: input.TreasuryId,
		EntityId:   input.EntityId,
	}

	// check if relations exists
	records := make(map[int]interface{})
	if input.TreasuryId != 0 {
		records[input.TreasuryId] = &models.Treasury{}
	}
	if input.EntityId != 0 {
		records[input.EntityId] = &models.Entity{}
	}
	if structName := errors.RecordNotFound(records); structName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "relation " + structName + " not found",
		})
		return
	}

	// create a new accounting_unit
	// preload relations
	if err := models.DB.Create(&accounting_unit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// return created accounting_unit
	if err := models.DB.Preload(clause.Associations).First(&accounting_unit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Initialize seriaizer
	serializer := serializers.AccountingUnitSerializer{
		C: c,
		AccountingUnit: accounting_unit,
	}

	c.JSON(http.StatusCreated, serializer.Response())
}

// update a accounting_unit
func UpdateAccountingUnit(c *gin.Context) {
	var accounting_unit = models.AccountingUnit{}
	// find accounting_unit by id
	if err := models.DB.Preload(clause.Associations).Where("id = ?", c.Param("id")).First(&accounting_unit).Error; err != nil {
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

	// check if relations exists
	records := make(map[int]interface{})
	if input.TreasuryId != 0 {
		records[input.TreasuryId] = &models.Treasury{}
	}
	if input.EntityId != 0 {
		records[input.EntityId] = &models.Entity{}
	}

	if structName := errors.RecordNotFound(records); structName != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "relation " + structName + " not found",
		})
		return
	}

	// update a accounting_unit
	models.DB.Model(&accounting_unit).Updates(models.AccountingUnit{
		NameEn:     input.NameEn,
		Status:     input.Status,
		TreasuryId: input.TreasuryId,
		EntityId:   input.EntityId,
	})

	// Initialize seriaizer
	serializer := serializers.AccountingUnitSerializer{
		C: c,
		AccountingUnit: accounting_unit,
	}

	c.JSON(http.StatusOK, serializer.Response())
}

// get unpaginated list of accounting_units
func GetUnpaginatedAccountingUnits(c *gin.Context) {
	var accounting_units []models.AccountingUnit
	if err := models.DB.Find(&accounting_units).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Initialize seriaizer 
	serializer := serializers.AccountingUnitsSerializer{
		C: c,
		AccountingUnits: accounting_units,
	}

	c.JSON(http.StatusOK, serializer.Response())
}
