package services

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"errors"
	"log"
)

type UserService interface {
	GetUserReviews(userID int) ([]models.Review, error)
	UpdateProfile(user *models.User) error
	RegisterUser(email, password, username string) (int, error)
	LoginUser(email, password string) (models.User, error)
}

type userService struct {
	rep *repository.Repository
}

func NewUserService(rep *repository.Repository) UserService {
	return &userService{rep: rep}
}

func (s *userService) GetUserReviews(userID int) ([]models.Review, error) {
	return s.rep.Review.GetUserReviews(userID)
}

func (s *userService) UpdateProfile(user *models.User) error {
	return s.rep.User.UpdateUserProfile(*user)
}

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

func (s *userService) LoginUser(email, password string) (models.User, error) {
	user, err := s.rep.User.GetUserByEmail(email)
	if err != nil {
		log.Printf("Ошибка получения пользователя по email %q: %v", email, err)
		return user, err
	}

	log.Printf("Пользователь найден: %s, хеш пароля: %s", user.Username, user.Password)

	if !utils.CheckPasswordHash(password, user.Password) {
		log.Printf("Сравнение хеша не прошло. Введённый пароль: %q, ожидаемый хеш: %s", password, user.Password)
		return user, errors.New("Неверный пароль")
	}

	return user, nil
}
