package db

import (
	"agrolumen/internal/config"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func GetDB() *sqlx.DB {
	once.Do(func() {
		config.LoadEnv()

		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.GetEnv("DB_USER"),
			config.GetEnv("DB_PASS"),
			config.GetEnv("DB_HOST"),
			config.GetEnv("DB_PORT"),
			config.GetEnv("DB_NAME"),
		)
		var err error
		db, err = sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Fatalf("Erro ao conectar no banco de dados: %v", err)
		}
	})
	return db
}
