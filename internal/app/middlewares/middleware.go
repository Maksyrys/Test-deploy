package app

import (
	"BookStore/internal/app/utils"
	"context"
	"log"
	"net/http"
	"strconv"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Паника: %v", err)
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
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
		// Если это предзапрос (preflight), завершаем обработку.
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *App) UserSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_user_id")
		if err == nil && cookie.Value != "" {
			// Предположим, в cookie мы храним прямое значение userID (строка)
			userID, err := strconv.Atoi(cookie.Value)
			if err == nil {
				user, err := a.rep.User.GetUserByID(userID)
				if err == nil {
					// Кладём в контекст
					ctx := context.WithValue(r.Context(), utils.UserContextKey, &user)
					r = r.WithContext(ctx)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
