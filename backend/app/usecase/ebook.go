package usecase

import (
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/mayukorin/ebook-merge/app/domain/repository"
)

type EbookUseCase struct {
	db *sqlx.DB
}

func NewEbookUseCase(db *sqlx.DB) *EbookUseCase {
	return &EbookUseCase{
		db: db,
	}
}

func (e *EbookUseCase) Index(userId int64) ([]model.Ebook, error) {
	return repository.AllEbooks(e.db, userId)
}
