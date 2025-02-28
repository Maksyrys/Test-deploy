package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	booksByCat, err := a.rep.Book.GetBooksGroupedByCategoryRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.templates.ExecuteTemplate(w, "index.html", booksByCat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) bookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Отсутствует параметр id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный параметр id", http.StatusBadRequest)
		return
	}

	book, err := a.rep.Book.GetBookByID(id)
	if err != nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}

	tmpl := NewTemplate("../../templates/book.html")
	err = tmpl.templates.ExecuteTemplate(w, "book.html", book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
