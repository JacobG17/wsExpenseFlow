package routes

import (
	"github.com/gin-gonic/gin"
	"Api/wsExpenseFlow/internal/handlers"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// Path para endpoint's de "User"
	userGroup := r.Group("/user")
	{
		userGroup.POST("/singUP", handlers.PostSingUpHandler)
	}
	//Ruta para dar de alta un Usuario
	return r
}