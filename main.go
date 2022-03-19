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
			"message": "Welcome to common financial management system",
		})
	})

	// get all treasuries
	r.GET("/treasuries", controllers.GetTreasuries)
	// create a new treasury
	r.POST("/treasuries", controllers.CreateTreasury)
	// update a treasury
	r.PUT("/treasuries/:id", controllers.UpdateTreasury)

	// get all accounting_units
	r.GET("/accounting_units", controllers.GetAccountingUnits)
	// create a new accounting_unit
	r.POST("/accounting_units", controllers.CreateAccountingUnit)
	// update a accounting_unit
	r.PUT("/accounting_units/:id", controllers.UpdateAccountingUnit)

	// get all entities
	r.GET("/entities", controllers.GetEntities)
	// create a new entity
	r.POST("/entities", controllers.CreateEntity)
	// update a entity
	r.PUT("/entities/:id", controllers.UpdateEntity)
	r.Run()
}
