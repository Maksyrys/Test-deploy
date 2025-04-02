package services

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"strconv"
)

type CatalogData struct {
	Categories           []models.Category
	SelectedCategoryID   int
	SelectedCategoryName string
	Books                []models.Book
}

type CatalogService interface {
	GetCatalogData(catIDStr string) (*CatalogData, error)
}

type catalogService struct {
	rep *repository.Repository
}

func NewCatalogService(rep *repository.Repository) CatalogService {
	return &catalogService{rep: rep}
}

func (s *catalogService) GetCatalogData(catIDStr string) (*CatalogData, error) {
	categories, err := s.rep.Book.GetAllCategories()
	if err != nil {
		return nil, err
	}

	data := &CatalogData{
		Categories: categories,
	}

	if len(categories) == 0 {
		return data, nil
	}

	var catID int
	if catIDStr != "" {
		if parsedID, err := strconv.Atoi(catIDStr); err == nil && parsedID > 0 {
			catID = parsedID
		}
	}

	if catID == 0 {
		catID = categories[0].ID
	}

	data.SelectedCategoryID = catID

	books, err := s.rep.Book.GetBooksByCategoryID(catID)
	if err != nil {
		return nil, err
	}
	data.Books = books

	for _, c := range categories {
		if c.ID == catID {
			data.SelectedCategoryName = c.Name
			break
		}
	}

	return data, nil
}
