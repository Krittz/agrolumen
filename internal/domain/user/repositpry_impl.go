package user

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}

// Create
func (r *repository) Create(u *User) error {
	query := `
	INSERT INTO users (nome, email, senha, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, u.Nome, u.Email, u.Senha, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

// Find By Email
func (r *repository) FindByEmail(email string) (*User, error) {
	query := `
	SELECT id, nome, email, senha, created_at, updated_at, deleted_at
	FROM users	
	WHERE email = ? AND deleted_at IS NULL 
	`
	var u User
	err := r.db.Get(&u, query, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Find By Id
func (r *repository) FindByID(id int64) (*User, error) {
	query := `
	SELECT id, nome, email, senha, created_at, updated_at, deleted_at
	FROM users
	WHERE id = ? AND deleted_at IS NULL
	`
	var u User

	err := r.db.Get(&u, query, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Update
func (r *repository) Update(u *User) error {
	query := `
	UPDATE users
	SET nome = ?, email = ?, senha = ?, updated_at = ?
	WHERE id = ? AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, u.Nome, u.Email, u.Senha, u.UpdatedAt, u.ID)
	return err
}

// Soft Delete
func (r *repository) Delete(id int64) error {
	query := `
	UPDATE users
	SET deleted_at = ?
	WHERE id = ? AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, time.Now(), id)
	return err

}
