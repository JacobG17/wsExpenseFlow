package main

import (
	"Api/wsExpenseFlow/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configurar las rutas
	routes.SetUpRouter()
 	
	r.Run("localhost:8080")
}