package main

import (
	"Api/wsExpenseFlow/internal/database"
	"Api/wsExpenseFlow/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Abrir conexiones de base de datos
	database.OpenConection()

	defer database.CloseDatabase()

	// Usar el router configurado en SetUpRouter()
	r := routes.SetUpRouter()

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
