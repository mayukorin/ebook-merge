package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/pkg/errors"
)

func AllEbooks(db *sqlx.DB, userId int64) ([]model.Ebook, error) {
	var ebooks []model.Ebook
	err := db.Select(&ebooks, "select eb.id, eb.title, es.id as \"e_book_service.id\", es.name as \"e_book_service.name\" from purchases as p join e_books as eb on p.e_book_id = eb.id join e_book_services as es on eb.e_book_service_id = es.id where p.user_id = ?", userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list ebook")
	}
	return ebooks, nil
}

func FindEbookByEbookServiceIDAndTitle(db *sqlx.DB, eBookServiceId int64, eBookTitle string) (*model.Ebook, error) {
	var ebook model.Ebook
	err := db.Get(&ebook, "select id, title, e_book_service_id as \"e_book_service.id\", created_at from e_books where e_book_service_id = ? and title = ?", eBookServiceId, eBookTitle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find ebook")
	}
	return &ebook, nil
}

func InsertEbook(db *sqlx.Tx, eBook *model.Ebook) (int64, error) {
	var id int64
	insertResult, err := db.Exec("insert into e_books (title, e_book_service_id) values (?, ?)", eBook.Title, eBook.EbookService.ID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert eBook")
	}
	id, err = insertResult.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "failed to select last insert id")
	}
	return id, nil
}
