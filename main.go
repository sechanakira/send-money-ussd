package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:changeit@tcp(127.0.0.1:3306)/ussd_menu?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to open database connection")
	}

	mysql, _ := db.DB()

	defer mysql.Close()

	db.AutoMigrate(&UssdSession{})
	db.AutoMigrate(&MenuText{})
	db.AutoMigrate(&ForeignCurrencyRates{})
}
