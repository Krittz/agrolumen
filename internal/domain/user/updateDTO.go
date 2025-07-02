package user

type UpdateDTO struct {
	Nome  *string `json:"nome,omitempty"`
	Email *string `json:"email,omitempty"`
	Senha *string `json:"senha,omitempty"`
}