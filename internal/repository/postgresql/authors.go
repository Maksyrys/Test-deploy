package postgresql

type Author struct {
	AuthorId  int
	Name      string
	Biography string
}

func NewAuthor() Author {
	return Author{}
}

func (a *Author) AddNewAuthor(id int, name string) error {
	return nil
}
