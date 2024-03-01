package postgres

import (
	"database/sql"
	"fmt"
	"goods/configs"

	_ "github.com/lib/pq"
)

func InitPostgre(cfg configs.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
	))

	return db, err
}
