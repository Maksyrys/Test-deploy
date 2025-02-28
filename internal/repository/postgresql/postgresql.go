package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	BDName   string
	SSLMode  string
}

func NewBD(config Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.Username, config.Password, config.BDName, config.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseBD(db *sql.DB) error {
	return db.Close()
}
