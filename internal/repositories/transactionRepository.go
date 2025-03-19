package repositories

import (
	"Api/wsExpenseFlow/internal/models"
	"gorm.io/gorm"
)

// TransactionRepository define el repositorio para manejar transacciones
type TransactionRepository struct {
	DB *gorm.DB
}

// NewTransaction crea una nueva transacción y la guarda en la base de datos
func (r *TransactionRepository) NewTransaction(transaction models.Transaction) (models.Transaction, error) {
	// Intentamos insertar la transacción en la base de datos
	if err := r.DB.Create(&transaction).Error; err != nil {
		// Si hay un error, lo devolvemos
		return models.Transaction{}, err
	}
	// Si todo va bien, devolvemos la transacción creada
	return transaction, nil
}