package app

import (
	"BookStore/internal/config"
	"BookStore/internal/repository"
	"BookStore/internal/repository/postgresql"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	rep *repository.Repository
}

func NewApp(rep *repository.Repository) *App {
	return &App{rep: rep}
}

func InitApp(db *sql.DB, cfg *config.Config) (*App, error) {
	rep := repository.NewRepository(db)
	app := NewApp(rep)

	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.Use(recoveryMiddleware)
	router.Use(corsMiddleware)
	router.Use(authMiddleware)

	fs := http.FileServer(http.Dir(cfg.Dir))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", app.indexHandler).Methods("GET")
	router.HandleFunc("/book/{id}", app.bookHandler).Methods("GET")

	log.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

	return app, nil
}

func Run() {
	cfg := config.NewConfig()

	db, err := InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresql.CloseBD(db)

	app, err := InitApp(db, cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app)
}

func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := postgresql.NewBD(postgresql.Config{
		Host:     cfg.BD.Host,
		Port:     cfg.BD.Port,
		Username: cfg.BD.Username,
		Password: cfg.BD.Password,
		BDName:   cfg.BD.BDName,
		SSLMode:  cfg.BD.SSLMode,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
