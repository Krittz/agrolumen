package user

type LoginDTO struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required"`
}
