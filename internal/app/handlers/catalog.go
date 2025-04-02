package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"net/http"
)

func (h *Handler) CatalogHandler(w http.ResponseWriter, r *http.Request) {
	catIDStr := r.URL.Query().Get("cat_id")
	catalogData, err := h.catalogService.GetCatalogData(catIDStr)
	if err != nil {
		appErr := utils.NewAppError("Ошибка при получении категорий", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	user := utils.GetCurrentUser(r)
	data := struct {
		CurrentUser          *models.User
		Categories           []models.Category
		SelectedCategoryID   int
		SelectedCategoryName string
		Books                []models.Book
	}{
		CurrentUser:          user,
		Categories:           catalogData.Categories,
		SelectedCategoryID:   catalogData.SelectedCategoryID,
		SelectedCategoryName: catalogData.SelectedCategoryName,
		Books:                catalogData.Books,
	}

	utils.Render(w, "../../templates/catalog.html", data)
}
