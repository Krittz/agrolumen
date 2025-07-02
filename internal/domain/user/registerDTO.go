package user

type SignUpDTO struct {
	Nome  string `json:"nome" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=8"`
}
