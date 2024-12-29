package services

import (
	"fmt"

	"github.com/TroJanBoi/CEPP-67/config"
	"github.com/TroJanBoi/CEPP-67/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.NewConfig().DB_HOST,
		config.NewConfig().DB_USER,
		config.NewConfig().DB_PASS,
		config.NewConfig().DB,
		config.NewConfig().DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// model table
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	return db, err
}

// func SetUpServer(app *fiber.App) {
// 	port := config.NewConfig().PORT
// 	log.Fatal(app.Listen(":" + port))
// }
