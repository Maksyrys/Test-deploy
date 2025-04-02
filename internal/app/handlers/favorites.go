package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	favorites, err := h.favoriteService.GetFavorites(user.UserId)
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения избранного", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	data := struct {
		CurrentUser *models.User
		Favorites   []models.Book
	}{
		CurrentUser: user,
		Favorites:   favorites,
	}
	utils.Render(w, "../../templates/favorites.html", data)
}

func (h *Handler) AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	if err := r.ParseForm(); err != nil {
		appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}
	bookIDStr := r.PostFormValue("book_id")
	if bookIDStr == "" {
		appErr := utils.NewAppError("Не указан идентификатор книги", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		appErr := utils.NewAppError("Неверный идентификатор книги", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}
	if err := h.favoriteService.AddFavorite(user.UserId, bookID); err != nil {
		appErr := utils.NewAppError("Ошибка при добавлении в избранное", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}

func (h *Handler) RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	if err := r.ParseForm(); err != nil {
		appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}
	bookIDStr := r.PostFormValue("book_id")
	if bookIDStr == "" {
		appErr := utils.NewAppError("Не указан идентификатор книги", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		appErr := utils.NewAppError("Неверный идентификатор книги", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}
	if err := h.favoriteService.RemoveFavorite(user.UserId, bookID); err != nil {
		appErr := utils.NewAppError("Ошибка при удалении из избранного", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}
