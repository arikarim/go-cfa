package controllers

import (
	"fmt"
	"github.com/arikarim/go-cfa/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

	c.JSON(200, treasuries)
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

	c.JSON(200, gin.H{
		"message": "success",
		"data":    treasury,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    treasury,
	})
}
