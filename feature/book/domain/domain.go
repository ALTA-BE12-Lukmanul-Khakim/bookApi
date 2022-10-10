package domain

type Basic struct {
	ID     uint
	Judul  string
	Author string
	IdUser uint
}

type Repository interface {
	Insert(newBook Basic) (Basic, error)
	Update(updatedBook Basic) (Basic, error)
	GetAll() ([]Basic, error)
	GetBook(ID uint) (Basic, error)
	Delete(ID uint) (Basic, error)
}

type Service interface {
	AddBook(newBook Basic) (Basic, error)
	EditBook(updatedBook Basic) (Basic, error)
	AllBook() ([]Basic, error)
	ThisBook(ID uint) (Basic, error)
	DeleteBook(ID uint) (Basic, error)
}
