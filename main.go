package main

import(
  "github.com/gin-gonic/gin"
	"github.com/arikarim/go-cfa/models"
	"github.com/arikarim/go-cfa/controllers"
)

func main() {
	r := gin.Default()

	// connect to database
	models.ConnectDatabase()

	// root of the project
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to common financial accounting units management system",
		})
	})

	// Treasury Group
	treasuryGroup := r.Group("v1/treasuries")
	{
		// get all treasuries
		treasuryGroup.GET("/", controllers.GetTreasuries)
		// get a treasury by id
		treasuryGroup.GET("/:id", controllers.GetTreasury)
		// create a new treasury
		treasuryGroup.POST("/", controllers.CreateTreasury)
		// update a treasury
		treasuryGroup.PUT("/:id", controllers.UpdateTreasury)
	}

	// accountingGroup Group
	accountingGroup := r.Group("v1/accounting_units")
	{
		// get all accounting_units
		accountingGroup.GET("/", controllers.GetAccountingUnits)
		// get a accounting_unit by id
		accountingGroup.GET("/:id", controllers.GetAccountingUnit)
		// create a new accounting_unit
		accountingGroup.POST("/", controllers.CreateAccountingUnit)
		// update a accounting_unit
		accountingGroup.PUT("/:id", controllers.UpdateAccountingUnit)
		// unpaginated list of accounting_units
		accountingGroup.GET("/all", controllers.GetUnpaginatedAccountingUnits)
	}

	// Entity Group
	entityGroup := r.Group("v1/entities")
	{
		// get all entities
		entityGroup.GET("/", controllers.GetEntities)
		// get a entity by id
		entityGroup.GET("/:id", controllers.GetEntity)
		// create a new entity
		entityGroup.POST("/", controllers.CreateEntity)
		// update a entity
		entityGroup.PUT("/:id", controllers.UpdateEntity)
	}
	r.Run()
}
