package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/pkg/errors"
)

func FindEbookServiceByName(db *sqlx.DB, name string) (*model.EbookService, error) {
	var ebookService model.EbookService
	err := db.Get(&ebookService, "select id, name, created_at from e_book_services where name = ?", name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find ebook_service")
	}
	return &ebookService, nil
}
