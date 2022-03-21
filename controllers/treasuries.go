package controllers

import (
	"fmt"
	"net/http"

	"github.com/arikarim/go-cfa/models"
	"github.com/arikarim/go-cfa/serializers"
	"github.com/gin-gonic/gin"
)

// get all treasuries
func GetTreasuries(c *gin.Context) {
	var treasuries []models.Treasury
	if models.DB.Find(&treasuries).Error != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	// Initialize serializer
	serializer := serializers.TreasuriesSerializer{
		C: c,
		Treasuries: treasuries,
	}

	c.JSON(200, serializer.Response())
}

// get an treasury by id
func GetTreasury(c *gin.Context) {
	var treasury = models.Treasury{}
	// find treasury by id
	if err := models.DB.Preload("AccountingUnits").Where("id = ?", c.Param("id")).First(&treasury).Error; err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Initialize serializer
	serializer := serializers.TreasurySerializer{
		C: c,
		Treasury: treasury,
	}
	// return success
	c.JSON(200, serializer.Response())
}

// create a new treasury
func CreateTreasury(c *gin.Context) {
	var input models.CreateTreasuryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Print("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// create a new treasury
	treasury := models.Treasury{
		NameEn: input.NameEn,
		Status: input.Status,
	}

	models.DB.Create(&treasury)

	// Initialize serializer
	serializer := serializers.TreasurySerializer{
		C: c,
		Treasury: treasury,
	}
	// return success
	c.JSON(http.StatusCreated, serializer.Response())
}

// UpdateTreasury
func UpdateTreasury(c *gin.Context) {
	var treasury = models.Treasury{}
	// find treasury by id
	if err := models.DB.Where("id = ?", c.Param("id")).First(&treasury).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	// validate input
	var input models.UpdateTreasuryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	fmt.Print("treasury: ", input)
	models.DB.Model(&treasury).Updates(models.Treasury{
		NameEn: input.NameEn,
		Status: input.Status,
	})

	// Initialize serializer
	serializer := serializers.TreasurySerializer{
		C: c,
		Treasury: treasury,
	}
	// return success
	c.JSON(http.StatusOK, serializer.Response())
}
