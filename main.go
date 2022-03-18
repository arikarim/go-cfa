package main

import(
	"github.com/gin-gonic/gin"
	"go.common-financial/models"
)

func main() {
	r := gin.Default()

	// root of the project
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to common financial management system",
		})
	})

	r.GET("/treasuries", models.GetTreasuries)

	r.Run()
}
