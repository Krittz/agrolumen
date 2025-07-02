package user

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(dto SignUpDTO) (*UserDTO, error)
	Login(dto LoginDTO) (*UserDTO, error)
	Me(userID int64) (*UserDTO, error)
	Update(userID int64, dto UpdateDTO) error
	Delete(userID int64) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

// Register

func (s *service) Register(dto SignUpDTO) (*UserDTO, error) {
	dto.Email = strings.ToLower(strings.TrimSpace(dto.Email))

	existing, _ := s.repo.FindByEmail(dto.Email)
	if existing != nil {
		return nil, errors.New("e-mail inválido")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Nome:      dto.Nome,
		Email:     dto.Email,
		Senha:     string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return &UserDTO{
		ID:    user.ID,
		Nome:  user.Nome,
		Email: user.Email,
	}, nil
}

// Login

func (s *service) Login(dto LoginDTO) (*UserDTO, error) {
	dto.Email = strings.ToLower(strings.TrimSpace(dto.Email))

	user, err := s.repo.FindByEmail(dto.Email)
	if err != nil || user == nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(dto.Senha)); err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	return &UserDTO{
		ID:    user.ID,
		Nome:  user.Nome,
		Email: user.Email,
	}, nil
}

// Me
func (s *service) Me(userID int64) (*UserDTO, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &UserDTO{
		ID:    user.ID,
		Nome:  user.Nome,
		Email: user.Email,
	}, nil
}

// Update

func (s *service) Update(userID int64, dto UpdateDTO) error {
	user, err := s.repo.FindByID(userID)
	if err != nil || user == nil {
		return errors.New("usuário não encontrado")
	}

	if dto.Nome != nil {
		user.Nome = *dto.Nome
	}

	if dto.Email != nil {
		user.Email = strings.ToLower(strings.TrimSpace(*dto.Email))
	}

	user.UpdatedAt = time.Now()

	if dto.Senha != nil && *dto.Senha != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Senha = string(hashedPassword)
	}

	return s.repo.Update(user)
}

// Delete
func (s *service) Delete(userID int64) error {
	_, err := s.repo.FindByID(userID)
	if err != nil {
		return errors.New("usuário não encontrado")
	}
	return s.repo.Delete(userID)
}
