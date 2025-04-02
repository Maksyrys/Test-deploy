package app

import (
	"BookStore/internal/app/middlewares"
	"BookStore/internal/config"
	"BookStore/internal/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Router struct {
	router *mux.Router
}

func NewRouter(rep *repository.Repository) *Router {
	return &Router{router: mux.NewRouter()}
}

func InitRouter(app *App, cfg *config.Config) {
	router := mux.NewRouter()

	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.RecoveryMiddleware)
	router.Use(middlewares.CorsMiddleware)
	router.Use(app.middle.UserSessionMiddleware)

	fs := http.FileServer(http.Dir(cfg.Dir))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", app.handler.IndexHandler).Methods("GET")
	router.HandleFunc("/book/{id}", app.handler.BookHandler).Methods("GET")

	router.HandleFunc("/register", app.handler.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", app.handler.LoginHandler).Methods("POST")

	router.HandleFunc("/profile", app.handler.ProfileHandler).Methods("GET")
	router.HandleFunc("/search", app.handler.SearchHandler).Methods("GET")

	router.HandleFunc("/cart", app.handler.CartHandler).Methods("GET")
	router.HandleFunc("/cart/add", app.handler.AddToCartHandler).Methods("POST")
	router.HandleFunc("/cart/remove", app.handler.RemoveFromCartHandler).Methods("POST")
	router.HandleFunc("/cart/checkout", app.handler.AddToCartHandler).Methods("POST")

	router.HandleFunc("/logout", app.handler.LogoutHandler).Methods("GET")
	router.HandleFunc("/cart/count", app.handler.CartCountHandler).Methods("GET")
	router.HandleFunc("/cart/remove/all", app.handler.RemoveAllFromCartHandler).Methods("POST")

	router.HandleFunc("/favorites", app.handler.FavoritesHandler).Methods("GET")
	router.HandleFunc("/favorites/add", app.handler.AddFavoriteHandler).Methods("POST")
	router.HandleFunc("/favorites/remove", app.handler.RemoveFavoriteHandler).Methods("POST")

	router.HandleFunc("/catalog", app.handler.CatalogHandler).Methods("GET")
	router.HandleFunc("/profile/edit", app.handler.EditProfileHandler).Methods("GET", "POST")
	router.HandleFunc("/book/{id}/reviews", app.handler.AddReviewHandler).Methods("POST")
	router.HandleFunc("/reviews/add", app.handler.CreateReviewHandler).Methods("POST")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("", app.handler.AdminDashboardHandler).Methods("GET")
	adminRouter.Use(middlewares.AdminMiddleware) // проверка, что пользователь admin

	adminRouter.HandleFunc("/book/add", app.handler.AdminAddBookHandler).Methods("POST")
	adminRouter.HandleFunc("/book/delete", app.handler.AdminDeleteBookHandler).Methods("POST")
	adminRouter.HandleFunc("/review/delete", app.handler.AdminDeleteReviewHandler).Methods("POST")

	log.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
