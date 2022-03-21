package controllers

import (
	"net/http"

	"github.com/arikarim/go-cfa/models"
	"github.com/arikarim/go-cfa/serializers"
	"github.com/gin-gonic/gin"
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

	serializer := serializers.EntitiesSerializer{
		C: c,
		Entities: entities,
	}

	// return success
	c.JSON(200, serializer.Response())
}

// get an entity by id
func GetEntity(c *gin.Context) {
	var entity = models.Entity{}
	// find entity by id
	// preload all accounting_units
	if err := models.DB.Preload("AccountingUnits").Where("id = ?", c.Param("id")).First(&entity).Error; err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Initialize serializer
	serializer := serializers.EntitySerializer{
		C: c, 
		Entity: entity,
	}
	// return success
	c.JSON(200, serializer.Response())
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

	// Initialize serializer
	serializer := serializers.EntitySerializer{
		C: c, 
		Entity: entity,
	}
	// return success
	c.JSON(200, serializer.Response())
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

	// Initialize serializer
	serializer := serializers.EntitySerializer{
		C: c, 
		Entity: entity,
	}
	// return success
	c.JSON(200, serializer.Response())
}
