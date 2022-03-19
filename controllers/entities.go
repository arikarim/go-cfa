package controllers

import (
	"github.com/arikarim/go-cfa/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get all entities
func GetEntities(c *gin.Context) {
	var entities []models.Entity
	if err := models.DB.Find(&entities).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// return success
	c.JSON(200, entities)
}

// create a new entity
func CreateEntity(c *gin.Context) {
	var input models.CreateEntityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		// return
		return
	}

	// create a new entity
	entity := models.Entity{
		NameEn: input.NameEn,
		Status: input.Status,
	}

	// Start a transaction
	tx := models.DB.Begin()
	// create a new entity
	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// commit transaction
	tx.Commit()

	// return success
	c.JSON(200, gin.H{
		"message": "success",
		"data":    entity,
	})
}

// update a entity
func UpdateEntity(c *gin.Context) {
	var entity = models.Entity{}
	// find entity by id
	if err := models.DB.Where("id = ?", c.Param("id")).First(&entity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	// validate input
	var input models.UpdateEntityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// update entity
	entity.NameEn = input.NameEn
	entity.Status = input.Status
	// update entity
	models.DB.Model(&entity).Updates(models.Entity{
		NameEn: input.NameEn,
		Status: input.Status,
	})

	// return success
	c.JSON(200, gin.H{
		"message": "success",
		"data":    entity,
	})
}