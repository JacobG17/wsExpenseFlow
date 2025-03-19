package handlers

import (
	"Api/wsExpenseFlow/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Funcion que maneja la peticion POST para dar de alta un usuario
func PostSignUpHandler(c *gin.Context) {
	var newUser models.User

	// BindJSON vincula el cuerpo de la solicitud con la estructura de models
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Solicitud Invalida",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messege": "Usuario recibido correctamente",
		"usuario": newUser,
	})
}
