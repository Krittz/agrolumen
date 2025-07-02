package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service Service
	jwtKey  []byte
}

func NewHandler(s Service, jwtKey string) *Handler {
	return &Handler{
		service: s,
		jwtKey:  []byte(jwtKey),
	}
}
func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/signup", h.SignUp)
	r.Post("/login", h.Login)

	r.Group(func(protected chi.Router) {
		protected.Use(h.AuthMiddleware)
		protected.Get("/me", h.Me)
		protected.Put("/update", h.Update)
		protected.Patch("/update", h.Update)
		protected.Delete("/delete", h.Delete)
	})
}
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var dto SignUpDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	user, err := h.service.Register(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.ID))

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var dto LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	user, err := h.service.Login(dto)
	if err != nil {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(h.jwtKey)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	user, err := h.service.Me(userID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var dto UpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	userID := userIDFromContext(r.Context())
	err := h.service.Update(userID, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	err := h.service.Delete(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("método de assinatura inválido")
			}
			return h.jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		userID := int64(claims["user_id"].(float64))
		ctx := contextWithUserID(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
