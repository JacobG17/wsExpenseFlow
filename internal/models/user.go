package models

// Definir la estructura de Go que representa los datos del usuario
type User struct {
	IDUsuario   int     `json:"id_usuario"`   // Nota: Usamos `json` tags para mapear con el JSON
	Nombre      string  `json:"nombre"`
	Apellido    string  `json:"apellido"`
	Email       string  `json:"email"`
	Contrasena  string  `json:"contrasena"`
	SaldoActual float64 `json:"saldo_actual"`
}