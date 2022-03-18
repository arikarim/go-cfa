package models

// import gorm "github.com/jinzhu/gorm"
// import postgres "github.com/jinzhu/gorm/dialects/postgres"
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=postgres dbname=common-financial-accounts sslmode=disable password=postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Treasury{})
}
