package services

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"errors"
)

// UserService описывает интерфейс бизнес-логики, связанной с пользователями.
type UserService interface {
	GetUserReviews(userID int) ([]models.Review, error)
	UpdateProfile(user *models.User) error
	RegisterUser(email, password, username string) (int, error)
	LoginUser(email, password string) (models.User, error)
}

// UserService – конкретная реализация UserService.
type userService struct {
	rep *repository.Repository
}

// NewUserService возвращает новый экземпляр userService.
func NewUserService(rep *repository.Repository) UserService {
	return &userService{rep: rep}
}

// GetUserReviews получает отзывы пользователя через репозиторий.
func (s *userService) GetUserReviews(userID int) ([]models.Review, error) {
	return s.rep.Review.GetUserReviews(userID)
}

// UpdateProfile обновляет профиль пользователя.
func (s *userService) UpdateProfile(user *models.User) error {
	return s.rep.User.UpdateUserProfile(*user)
}

// RegisterUser регистрирует нового пользователя: хеширует пароль и создаёт запись.
func (s *userService) RegisterUser(email, password, username string) (int, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return 0, err
	}

	user := models.User{
		Email:    email,
		Password: hashedPassword,
		Username: username,
	}

	return s.rep.User.CreateUser(user)
}

// LoginUser выполняет аутентификацию: ищет пользователя по email и проверяет пароль.
func (s *userService) LoginUser(email, password string) (models.User, error) {
	user, err := s.rep.User.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return user, errors.New("Неверный пароль")
	}

	return user, nil
}
