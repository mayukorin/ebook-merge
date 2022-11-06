package model

import (
	"time"

	"github.com/mayukorin/ebook-merge/swagger/generated_swagger"
)

type Ebook struct {
	ID           int64        `db:"id"`
	Title        string       `db:"title"`
	EbookService EbookService `db:"e_book_service"`
	CreatedAt    time.Time    `db:"created_at"`
}

func (e *Ebook) SwaggerModel() *generated_swagger.Ebook {
	return &generated_swagger.Ebook{
		ID:           e.ID,
		Title:        e.Title,
		EbookService: e.EbookService.SwaggerModel(),
	}
}
