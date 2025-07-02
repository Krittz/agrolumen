package main

import (
	"agrolumen/internal/config"
	"agrolumen/internal/db"
	"agrolumen/internal/domain/user"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//Carrega variáveis de ambiente
	config.LoadEnv()

	// Lê valores
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET nõ definido no .env")
	}

	// Conecta ao banco via singleton
	dbConn := db.GetDB()

	// Injeto dependências do domínio de usuário

	userRepo := user.NewRepository(dbConn)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService, jwtSecret)

	// Define router com middlewares
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Registra rotas
	r.Route("/api", func(api chi.Router) {
		api.Route("/users", userHandler.RegisterRoutes)
		//api.Route("/animals, animalHandler.RegisterRoutes") <- em breve
	})

	// Inicia servidor
	log.Printf("Servidor rodando em http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

}
