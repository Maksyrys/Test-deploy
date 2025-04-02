package handlers

import (
	"BookStore/internal/app/services"
	"BookStore/internal/repository"
	"github.com/gorilla/sessions"
)

type Handler struct {
	rep             *repository.Repository
	userService     services.UserService
	bookService     services.BookService
	cartService     services.CartService
	catalogService  services.CatalogService
	favoriteService services.FavoriteService
	reviewService   services.ReviewService
	store           sessions.Store
}

func NewHandler(rep *repository.Repository,
	userService services.UserService,
	bookService services.BookService,
	cartService services.CartService,
	catalogService services.CatalogService,
	favoriteService services.FavoriteService,
	reviewService services.ReviewService,
	store sessions.Store,
) *Handler {
	return &Handler{rep: rep,
		userService:     userService,
		bookService:     bookService,
		cartService:     cartService,
		catalogService:  catalogService,
		favoriteService: favoriteService,
		reviewService:   reviewService,
		store:           store,
	}
}
