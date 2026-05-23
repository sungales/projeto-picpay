package models

type Usuario struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Balance  int64  `json:"balance"`
}
