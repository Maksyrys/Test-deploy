package handlers

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/models"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	favoritesCount, err := h.favoriteService.GetFavoritesCount(user.UserId)
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения кол-ва избранных", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	cartCount, err := h.cartService.GetCartCount(user.UserId)
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения кол-ва книг в корзине", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	reviews, err := h.userService.GetUserReviews(user.UserId)
	if err != nil {
		appErr := utils.NewAppError("Ошибка получения отзывов", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	data := struct {
		CurrentUser    *models.User
		FavoritesCount int
		CartCount      int
		UserReviews    []models.Review
	}{
		CurrentUser:    user,
		UserReviews:    reviews,
		FavoritesCount: favoritesCount,
		CartCount:      cartCount,
	}

	utils.Render(w, "../../templates/profile.html", data)
}

func (h *Handler) EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		appErr := utils.NewAppError("Пользователь не авторизован", http.StatusUnauthorized, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
			utils.RespondWithError(w, appErr)
			return
		}

		user.Username = r.FormValue("username")
		user.Firstname = r.FormValue("firstname")
		user.Lastname = r.FormValue("lastname")
		user.Phone = r.FormValue("phone")

		if err := h.userService.UpdateProfile(user); err != nil {
			appErr := utils.NewAppError("Ошибка обновления профиля", http.StatusInternalServerError, err)
			utils.RespondWithError(w, appErr)
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

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	email := r.PostFormValue("reg-email")
	password := r.PostFormValue("reg-password")
	username := r.PostFormValue("reg-username")

	if email == "" || username == "" || password == "" {
		appErr := utils.NewAppError("Необходимо заполнить все поля", http.StatusBadRequest, nil)
		utils.RespondWithError(w, appErr)
		return
	}

	newUserID, err := h.userService.RegisterUser(email, password, username)
	if err != nil {
		appErr := utils.NewAppError("Ошибка при создании пользователя", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	session, err := h.store.Get(r, "session")
	if err != nil {
		appErr := utils.NewAppError("Не удалось получить сессию", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	session.Values["user_id"] = newUserID

	if err := session.Save(r, w); err != nil {
		appErr := utils.NewAppError("Не удалось сохранить сессию", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		appErr := utils.NewAppError("Ошибка парсинга формы", http.StatusBadRequest, err)
		utils.RespondWithError(w, appErr)
		return
	}

	email := r.FormValue("login-email")
	password := r.FormValue("login-password")
	// Здесь можно добавить структурированное логирование попытки входа
	utils.Logger.Info("Попытка входа", zap.String("email", email))
	user, err := h.userService.LoginUser(email, password)
	if err != nil {
		appErr := utils.NewAppError("Пользователь не найден или неверный пароль", http.StatusUnauthorized, err)
		utils.RespondWithError(w, appErr)
		return
	}

	// Создаём / получаем сессию
	session, err := h.store.Get(r, "session")
	if err != nil {
		appErr := utils.NewAppError("Ошибка при получении сессии", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}

	// Сохраняем userID в сессии
	session.Values["user_id"] = user.UserId

	// Дополнительно можете настроить время жизни, флаги безопасности и т.д.
	// Пример:
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 8, // 8 часов
		HttpOnly: true,
		// Secure:   true,       // Включайте, если используете https
	}

	// Сохраняем изменения в cookie
	if err = session.Save(r, w); err != nil {
		appErr := utils.NewAppError("Не удалось сохранить сессию", http.StatusInternalServerError, err)
		utils.RespondWithError(w, appErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"success": true}`))
}

func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := h.store.Get(r, "session")
	session.Options.MaxAge = -1 // Сбрасываем сессию
	_ = session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
