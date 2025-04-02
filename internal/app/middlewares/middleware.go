package middlewares

import (
	"BookStore/internal/app/utils"
	"BookStore/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"net/http"
)

type Middleware struct {
	rep   *repository.Repository
	store sessions.Store
}

func NewMiddleware(rep *repository.Repository, store sessions.Store) *Middleware {
	return &Middleware{rep: rep, store: store}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Info("Запрос",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				var err error
				if recErr, ok := rec.(error); ok {
					err = recErr
				} else {
					err = fmt.Errorf("%v", rec)
				}
				utils.Logger.Error("Паника в обработчике", zap.Error(err))
				appErr := utils.NewAppError("Внутренняя ошибка сервера", http.StatusInternalServerError, err)
				utils.RespondWithError(w, appErr)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Если это preflight-запрос, завершаем обработку.
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) UserSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := m.store.Get(r, "session")
		if err != nil {
			utils.Logger.Warn("Не удалось получить сессию", zap.Error(err))
			next.ServeHTTP(w, r)
			return
		}

		userID, ok := session.Values["user_id"].(int)
		if ok && userID > 0 {
			user, err := m.rep.User.GetUserByID(userID)
			if err == nil {
				utils.Logger.Info("Пользователь авторизован",
					zap.String("username", user.Username),
					zap.Int("userID", user.UserId),
					zap.String("role", user.Role),
				)

				ctx := context.WithValue(r.Context(), utils.UserContextKey, &user)
				r = r.WithContext(ctx)
			} else {
				utils.Logger.Warn("Не удалось получить пользователя по ID", zap.Int("userID", userID), zap.Error(err))
			}
		}
		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := utils.GetCurrentUser(r)
		if user == nil || user.Role != "admin" {
			appErr := utils.NewAppError("Доступ запрещён", http.StatusForbidden, errors.New("user is not admin"))
			utils.RespondWithError(w, appErr)
			return
		}
		next.ServeHTTP(w, r)
	})
}
