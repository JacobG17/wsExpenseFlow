package routes

import (
	"Api/wsExpenseFlow/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetUpRouter configura las rutas de la API
func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// Grupo de rutas para usuarios
	userGroup := r.Group("/user")
	{
		userGroup.POST("/sign-up", handlers.PostSignUpHandler) // Se descomenta si se usa
	}

	// Grupo de rutas para transacciones
	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.POST("/new", handlers.PostNewTransaction) // Ruta más RESTful
	}

	return r
}
