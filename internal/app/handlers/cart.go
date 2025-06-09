package handlers

import (
	"BookStore/internal/app/services"
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) CartHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	details, err := h.cartService.GetCartDetails(user.UserId)
	if err != nil {
		appErr := utils.NewAppError("Ошибка при получении корзины", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	data := struct {
		CurrentUser   *models.User
		Items         []services.CartItemDetail
		GrandTotal    float64
		TotalQuantity int
	}{
		CurrentUser:   user,
		Items:         details.Items,
		GrandTotal:    details.GrandTotal,
		TotalQuantity: details.TotalQuantity,
	}

	utils.Render(w, "templates/cart.html", data)
}

func (h *Handler) AddToCartHandler(w http.ResponseWriter, r *http.Request) {
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
	quantityStr := r.PostFormValue("quantity")
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
	quantity := 1
	if quantityStr != "" {
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			appErr := utils.NewAppError("Неверное количество", http.StatusBadRequest, err)
			utils.RespondWithError(w, appErr)
			return
		}
	}

	if err = h.cartService.AddItem(user.UserId, bookID, quantity); err != nil {
		appErr := utils.NewAppError("Ошибка при добавлении в корзину", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}

func (h *Handler) RemoveFromCartHandler(w http.ResponseWriter, r *http.Request) {
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

	if err = h.cartService.RemoveOneItem(user.UserId, bookID); err != nil {
		appErr := utils.NewAppError("Ошибка при удалении из корзины", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}

func (h *Handler) CartCountHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	totalCount, err := h.cartService.GetCartCount(user.UserId)
	if err != nil {
		appErr := utils.NewAppError("Ошибка при получении количества товаров в корзине", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf(`{"count": %d}`, totalCount)))
}

func (h *Handler) RemoveAllFromCartHandler(w http.ResponseWriter, r *http.Request) {
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

	if err = h.cartService.RemoveAllItems(user.UserId, bookID); err != nil {
		appErr := utils.NewAppError("Ошибка при удалении из корзины", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}
