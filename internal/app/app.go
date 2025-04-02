package app

import (
	"BookStore/internal/app/handlers"
	"BookStore/internal/app/middlewares"
	"BookStore/internal/app/services"
	"BookStore/internal/app/utils"
	"BookStore/internal/config"
	"BookStore/internal/repository"
	"BookStore/internal/repository/postgresql"
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	"log"
)

type App struct {
	rep     *repository.Repository
	handler *handlers.Handler
	middle  *middlewares.Middleware
	cfg     *config.Config
	Store   sessions.Store
}

func NewApp(rep *repository.Repository, cfg *config.Config) *App {
	userService := services.NewUserService(rep)
	bookService := services.NewBookService(rep)
	cartService := services.NewCartService(rep)
	catalogService := services.NewCatalogService(rep)
	favoriteService := services.NewFavoriteService(rep)
	reviewService := services.NewReviewService(rep)
	store := sessions.NewCookieStore([]byte(cfg.Key))

	return &App{rep: rep,
		handler: handlers.NewHandler(rep, userService, bookService, cartService, catalogService, favoriteService, reviewService, store),
		middle:  middlewares.NewMiddleware(rep, store),
		cfg:     cfg,
		Store:   store,
	}
}

func InitApp(db *sql.DB, cfg *config.Config) (*App, error) {
	rep := repository.NewRepository(db)
	app := NewApp(rep, cfg)

	InitRouter(app, cfg)

	return app, nil
}

func Run() {

	if err := utils.InitLogger(); err != nil {
		panic("Не удалось инициализировать логгер: " + err.Error())
	}
	defer utils.Logger.Sync()

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
