package utils

import (
	"BookStore/internal/models"
	"bytes"
	"embed"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"path/filepath"
)

type contextKey string

const UserContextKey = contextKey("currentUser")

var templatesFS embed.FS

func GetCurrentUser(r *http.Request) *models.User {
	val := r.Context().Value(UserContextKey)
	if val == nil {
		return nil
	}
	user, ok := val.(*models.User)
	if !ok {
		return nil
	}
	return user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Render(w http.ResponseWriter, templateFile string, data interface{}) {
	// Извлекаем только имя файла (например "admin.html")
	filename := filepath.Base(templateFile)
	// Собираем путь внутри embed.FS — "templates/admin.html"
	tplPath := filepath.Join("templates", filename)

	// Парсим layout + все общие элементы + нужный шаблон
	tmpl := template.Must(template.New("layout.html").ParseFS(templatesFS,
		"templates/layout.html",
		"templates/header.html",
		"templates/modal.html",
		"templates/footer.html",
		tplPath,
	))

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "layout", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
