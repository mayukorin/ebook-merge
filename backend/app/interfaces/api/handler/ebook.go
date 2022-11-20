package handler

import (
	"net/http"

	"github.com/mayukorin/ebook-merge/app/interfaces/api/httputil"
	"github.com/mayukorin/ebook-merge/app/usecase"
	"github.com/mayukorin/ebook-merge/swagger/generated_swagger"
)

type EbookHandler struct {
	ebookUseCase *usecase.EbookUseCase
}

func NewEbookHandler(ebookUseCase *usecase.EbookUseCase) *EbookHandler {
	return &EbookHandler{ebookUseCase: ebookUseCase}
}

func (e *EbookHandler) Index(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	ebooks, err := e.ebookUseCase.Index(user.ID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	var resEbooks []*generated_swagger.Ebook
	for _, e := range ebooks {
		resEbooks = append(resEbooks, e.SwaggerModel())
	}
	return http.StatusOK, generated_swagger.ListEbook{Ebooks: resEbooks}, nil
}

func (e *EbookHandler) ScanAllEbooksFromGmail(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	err = e.ebookUseCase.ScanAllEbooksFromGmail(user.ID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, nil, nil

}

func (e *EbookHandler) ScanTestEbooksFromGmail(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	err := e.ebookUseCase.TestScanKindleEbooksFromGmail(47)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, nil, nil

}
