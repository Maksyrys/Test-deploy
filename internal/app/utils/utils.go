package utils

import (
	"BookStore/internal/models"
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

func SetUserSessionCookie(w http.ResponseWriter, userID int) {
	cookie := http.Cookie{
		Name:     "session_user_id",
		Value:    strconv.Itoa(userID),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func Render(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(
		"../../templates/layout.html",
		"../../templates/header.html",
		"../../templates/modal.html",
		"../../templates/footer.html",
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

func SaveUploadedFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	fileExt := filepath.Ext(header.Filename)
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + fileExt

	filePath := filepath.Join("static", "uploads", fileName)
	dest, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	if err != nil {
		return "", err
	}

	// Исправляем слеши:
	webFilePath := "/" + filepath.ToSlash(filePath)

	return webFilePath, nil
}
