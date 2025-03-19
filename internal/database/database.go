package database

import (
	"Api/wsExpenseFlow/config"
	"Api/wsExpenseFlow/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConection() {
	dsn := config.GetConectionString()
	var err error
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	DB.AutoMigrate(&models.Transaction{})
}

func CloseDatabase() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error al obtener acceso a la base de datos:", err)
	}
	sqlDB.Close()
}
