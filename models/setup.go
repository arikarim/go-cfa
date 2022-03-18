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
	dsn := "host=localhost user=postgres dbname=postgres password=passowrd sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Treasury{})
}
