package utils

import (
	"BookStore/internal/models"
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type contextKey string

const UserContextKey = contextKey("currentUser")

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
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/header.html",
		"templates/modal.html",
		"templates/footer.html",
		templateFile,
	))

	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = buf.WriteTo(w)
}
