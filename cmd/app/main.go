package main

import (
	"Api/wsExpenseFlow/internal/database"
	"Api/wsExpenseFlow/internal/routes"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Obtener el directorio actual de trabajo
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error al obtener el directorio de trabajo:", err)
	}

	// Obtener la ruta del directorio raíz del proyecto
	projectRoot := filepath.Join(workDir, "../..")

	// Cargar variables de entorno desde la raíz del proyecto
	if err := godotenv.Load(filepath.Join(projectRoot, ".env")); err != nil {
		log.Fatal("Error al cargar el archivo .env:", err)
	}

	// Abrir conexiones de base de datos
	database.OpenConection()

	defer database.CloseDatabase()

	// Usar el router configurado en SetUpRouter()
	r := routes.SetUpRouter()

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
