package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/mayukorin/ebook-merge/app/interfaces/api/middleware"
	"github.com/rs/cors"
)

func NewRouter() *mux.Router {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "Authorization"},
	})
	commonChain := alice.New(c.Handler, middleware.RecoverMiddleware)
	r := mux.NewRouter()
	sr := r.PathPrefix("/v1").Subrouter()
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/ping").Handler(commonChain.Then(AppHandler{ping}))

	return r
}

func ping(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, "pong", nil
}
