package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mayukorin/ebook-merge/app/interfaces/api/httputil"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occured!")
				httputil.RespondErrByJSON(w, r, http.StatusInternalServerError, errors.New("panic occured"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
