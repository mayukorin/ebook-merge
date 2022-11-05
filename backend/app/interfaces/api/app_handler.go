package api

import (
	"net/http"

	"github.com/mayukorin/ebook-merge/app/interfaces/api/httputil"
)

type AppHandler struct {
	h func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, res, err := a.h(w, r)

	if err != nil {
		httputil.RespondErrByJSON(w, r, statusCode, err)
		return
	}

	httputil.RespondByJSON(w, r, statusCode, res)

	return

}
