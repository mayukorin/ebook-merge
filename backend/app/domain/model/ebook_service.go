package model

import (
	"time"

	"github.com/mayukorin/ebook-merge/swagger/generated_swagger"
)

type EbookService struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func (e *EbookService) SwaggerModel() *generated_swagger.EbookService {
	return &generated_swagger.EbookService{
		ID:   e.ID,
		Name: e.Name,
	}
}
