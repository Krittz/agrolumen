package user

import "time"

type User struct {
	ID        int64      `db:"id" json:"id"`
	Nome      string     `db:"nome" json:"nome"`
	Email     string     `db:"email" json:"email"`
	Senha     string     `db:"senha" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
