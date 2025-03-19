package handlers

import (
	"Api/wsExpenseFlow/internal/models"
	"net/http"

	"Api/wsExpenseFlow/internal/database"

	"github.com/gin-gonic/gin"
)

func PostNewTransaction(c *gin.Context) {
	var newTransaction models.Transaction

	// Deberías enlazar los datos de la solicitud (request) al modelo
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input transaction", "error": err.Error()})
		return
	}
	// Ahora vamos a intentar guardar la transacción en la base de datos
	database.DB.Create(&newTransaction)

	// Si la transacción se guarda correctamente, podemos devolver la respuesta
	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaction saved successfully",
	})
}
