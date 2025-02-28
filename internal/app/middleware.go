package app

import (
	"log"
	"net/http"
)

// loggingMiddleware выводит информацию о каждом запросе.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// recoveryMiddleware перехватывает панику и возвращает ошибку 500.
func recoveryMiddleware(next http.Handler) http.Handler {
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

// corsMiddleware добавляет заголовки для поддержки CORS.
func corsMiddleware(next http.Handler) http.Handler {
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

// authMiddleware проверяет наличие заголовка Authorization (пример простой проверки).
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Для демонстрации: если заголовок отсутствует, можно вернуть ошибку или пропустить.
		// Здесь пропускаем запрос, но можно добавить реальную логику аутентификации.
		token := r.Header.Get("Authorization")
		if token == "" {
			// Например, можно разрешить доступ к публичным ресурсам.
			// Или, если требуется авторизация, вернуть ошибку:
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			// return
		}
		next.ServeHTTP(w, r)
	})
}
