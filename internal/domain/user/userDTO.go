package user

type UserDTO struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
