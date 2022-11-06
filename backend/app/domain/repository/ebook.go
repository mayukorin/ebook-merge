package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/pkg/errors"
)

func AllEbooks(db *sqlx.DB, userId int64) ([]model.Ebook, error) {
	var ebooks []model.Ebook
	err := db.Select(&ebooks, "select eb.id, eb.title, es.id as \"e_book_service.id\", es.name as \"e_book_service.name\" from purchases as p join e_books as eb on p.e_book_id = eb.id join e_book_services as es on eb.e_book_service_id = es.id")
	if err != nil {
		return nil, errors.Wrap(err, "failed to list ebook")
	}
	return ebooks, nil
}
