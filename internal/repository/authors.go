package repository

type AuthorsRepository interface {
	AddNewAuthor(id int, name string) error
}
