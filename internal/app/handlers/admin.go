package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil || user.Role != "admin" {
		appErr := utils.NewAppError("Доступ запрещён", http.StatusForbidden, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	books := h.bookService.GetAllBooks()
	reviews, err := h.reviewService.GetAllReviews()
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения отзывов", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	categories, err := h.rep.Book.GetAllCategories()
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения категорий", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	data := struct {
		CurrentUser *models.User
		Books       []models.Book
		Reviews     []models.Review
		Categories  []models.Category
	}{
		CurrentUser: user,
		Books:       books,
		Reviews:     reviews,
		Categories:  categories,
	}
	utils.Render(w, "./templates/admin.html", data)
}

func (h *Handler) AdminAddBookHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil || user.Role != "admin" {
		appErr := utils.NewAppError("Доступ запрещён", http.StatusForbidden, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	if err := r.ParseForm(); err != nil {
		appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	category := r.FormValue("category")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	description := r.FormValue("description")
	detailedDescription := r.FormValue("detailed_description")

	book := models.Book{
		Title:               title,
		Author:              author,
		Category:            category,
		Price:               price,
		Description:         description,
		DetailedDescription: detailedDescription,
	}

	if err := h.bookService.AddBook(book); err != nil {
		appErr := utils.NewAppError("Ошибка при добавлении книги", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *Handler) AdminDeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil || user.Role != "admin" {
		appErr := utils.NewAppError("Доступ запрещён", http.StatusForbidden, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	idStr := r.FormValue("id")
	if idStr == "" {
		appErr := utils.NewAppError("Не указан id книги", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appErr := utils.NewAppError("Некорректный id", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	if err := h.bookService.DeleteBook(id); err != nil {
		appErr := utils.NewAppError("Ошибка удаления книги", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *Handler) AdminDeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil || user.Role != "admin" {
		appErr := utils.NewAppError("Доступ запрещён", http.StatusForbidden, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	reviewIDStr := r.FormValue("id")
	if reviewIDStr == "" {
		appErr := utils.NewAppError("Не указан id отзыва", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		appErr := utils.NewAppError("Некорректный id отзыва", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	if err := h.reviewService.DeleteReview(reviewID); err != nil {
		appErr := utils.NewAppError("Ошибка удаления отзыва", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (h *Handler) AdminAddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)

	if user == nil || user.Role != "admin" {
		appErr := utils.NewAppError("Доступ запрещё", http.StatusForbidden, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	if err := r.ParseForm(); err != nil {
		appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	categoryName := r.FormValue("categoryName")

	category := models.Category{Name: categoryName}

	if err := h.bookService.CreateCategory(category); err != nil {
		appErr := utils.NewAppError("Ошибка при добавлении категории", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
}
