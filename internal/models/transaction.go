package models

type Transaction struct {
	IDUser                 int     `json:"id_user"`
	Amount                 float64 `json:"amount"`
	Description            string  `json:"description"`
	TypeTransaction        string  `json:"type_transaction"`
	Category               string  `json:"category"`
	PostTransactionBalance float64 `json:"post_transaction_balance"`
	Date                   string  `json:"date"`
}
