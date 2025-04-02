package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	randomBooks, categories, err := h.bookService.GetIndexData(12)
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения данных для главной страницы", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	user := utils.GetCurrentUser(r)
	data := struct {
		RandomBooks []models.Book
		Categories  []models.Category
		CurrentUser *models.User
	}{
		RandomBooks: randomBooks,
		Categories:  categories,
		CurrentUser: user,
	}

	utils.Render(w, "../../templates/index.html", data)
}

func (h *Handler) BookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		appErr := utils.NewAppError("Отсутствует параметр id", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		appErr := utils.NewAppError("Неверный параметр id", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	user := utils.GetCurrentUser(r)
	bookDetails, err := h.bookService.GetBookDetails(id, user)
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения данных о книге", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	data := struct {
		BooksByAuthor    map[string][]models.Book
		AuthorBooks      []models.Book
		Book             models.Book
		CurrentUser      *models.User
		InCart           bool
		IsFavorite       bool
		BookReviews      []models.Review
		UserReviewExists bool
	}{
		BooksByAuthor:    bookDetails.BooksByAuthor,
		AuthorBooks:      bookDetails.AuthorBooks,
		Book:             bookDetails.Book,
		CurrentUser:      user,
		InCart:           bookDetails.InCart,
		IsFavorite:       bookDetails.IsFavorite,
		BookReviews:      bookDetails.BookReviews,
		UserReviewExists: bookDetails.UserReviewExists,
	}

	utils.Render(w, "../../templates/book.html", data)
}

func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	user := utils.GetCurrentUser(r)

	if query == "" {
		data := struct {
			CurrentUser *models.User
			Query       string
			Books       []models.Book
		}{
			CurrentUser: user,
			Query:       "",
			Books:       nil,
		}
		utils.Render(w, "../../templates/search.html", data)
		return
	}

	books, err := h.bookService.SearchBooks(query)
	if err != nil {
		appErr := utils.NewAppError("Ошибка поиска", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	data := struct {
		CurrentUser *models.User
		Query       string
		Books       []models.Book
	}{
		CurrentUser: user,
		Query:       query,
		Books:       books,
	}

	utils.Render(w, "../../templates/search.html", data)
}
