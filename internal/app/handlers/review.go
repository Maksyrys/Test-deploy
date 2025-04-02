package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) AddReviewHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		appErr := utils.NewAppError("Некорректный id книги", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		appErr := utils.NewAppError("Некорректный рейтинг", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	review := models.Review{
		UserID: user.UserId,
		BookID: bookID,
		Rating: rating,
	}

	if err := h.reviewService.CreateReview(review); err != nil {
		appErr := utils.NewAppError("Ошибка при создании отзыва", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/book/%d", bookID), http.StatusSeeOther)
}

func (h *Handler) CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	bookID, err := strconv.Atoi(r.FormValue("book_id"))
	if err != nil {
		appErr := utils.NewAppError("Некорректный book_id", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil || rating < 1 || rating > 5 {
		appErr := utils.NewAppError("Некорректный рейтинг", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	comment := r.FormValue("comment")
	review := models.Review{
		UserID:  user.UserId,
		BookID:  bookID,
		Rating:  rating,
		Comment: comment,
	}

	if err := h.reviewService.CreateReview(review); err != nil {
		appErr := utils.NewAppError("Ошибка сохранения отзыва", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
