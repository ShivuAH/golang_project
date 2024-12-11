package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sever/model"

	"github.com/google/uuid"
)

type BookRepository struct {
	con *sql.DB
}

func NewBookRepository(con *sql.DB) *BookRepository {
	return &BookRepository{
		con,
	}
}

func (br *BookRepository) book(r *http.Request) (string, error) {
	var newBook model.BookModel
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		return "", err
	}

	newBook.BookId = uuid.NewString()

	_ , err = br.con.Exec("INSERT INTO books(book_id , book_name, category_id, total_books)VALUES($1,$2,$3,$4)", newBook.BookId, newBook.BookName, newBook.CategoryId, newBook.TotalCopies)
	if err !=nil {
		return "Updated successfully",err
	}
	return "Updated successfully",nil
}