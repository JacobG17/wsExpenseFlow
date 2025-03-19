package main

import (
	"Api/wsExpenseFlow/internal/database"
	"Api/wsExpenseFlow/internal/routes"
)

func main() {
	// Abrir conexiones de base de datos
	database.OpenConection()

	//defer database.CloseDatabase()

	// Usar el router configurado en SetUpRouter()
	r := routes.SetUpRouter()

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
