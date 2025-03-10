package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"log"
	"net/http"
)

// ProfileHandler получает профиль пользователя и его отзывы через сервис.
func (h *Handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	reviews, err := h.userService.GetUserReviews(user.UserId)
	if err != nil {
		http.Error(w, "Ошибка получения отзывов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		CurrentUser *models.User
		UserReviews []models.Review
	}{
		CurrentUser: user,
		UserReviews: reviews,
	}

	utils.Render(w, "../../templates/profile.html", data)
}

// EditProfileHandler обновляет профиль пользователя через сервис.
func (h *Handler) EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
			return
		}

		user.Username = r.FormValue("username")
		user.Firstname = r.FormValue("firstname")
		user.Lastname = r.FormValue("lastname")
		user.Phone = r.FormValue("phone")

		if err := h.userService.UpdateProfile(user); err != nil {
			http.Error(w, "Ошибка обновления профиля: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	data := struct {
		CurrentUser *models.User
	}{
		CurrentUser: user,
	}
	utils.Render(w, "../../templates/edit_profile.html", data)
}

// RegisterHandler регистрирует нового пользователя через сервис.
func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}

	email := r.PostFormValue("reg-email")
	password := r.PostFormValue("reg-password")
	username := r.PostFormValue("reg-username")

	if email == "" || username == "" || password == "" {
		http.Error(w, "Необходимо заполнить все поля!", http.StatusBadRequest)
		return
	}

	newUserID, err := h.userService.RegisterUser(email, password, username)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Устанавливаем cookie с идентификатором нового пользователя
	utils.SetUserSessionCookie(w, newUserID)

	// Возвращаем JSON-ответ вместо редиректа
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

// LoginHandler выполняет вход пользователя через сервис аутентификации.
func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}

	email := r.FormValue("login-email")
	password := r.FormValue("login-password")
	log.Printf("Попытка входа: email = %q", email)

	user, err := h.userService.LoginUser(email, password)
	if err != nil {
		http.Error(w, "Пользователь не найден или неверный пароль", http.StatusUnauthorized)
		return
	}

	utils.SetUserSessionCookie(w, user.UserId)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

// LogoutHandler выполняет выход пользователя (очистка cookie).
func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "session_user_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}
