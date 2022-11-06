package api

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/mayukorin/ebook-merge/app/interfaces/api/handler"
	"github.com/mayukorin/ebook-merge/app/interfaces/api/middleware"
	"github.com/mayukorin/ebook-merge/app/usecase"
	"github.com/rs/cors"
)

func NewRouter(db *sqlx.DB, firebaseClient *auth.Client) *mux.Router {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "Authorization"},
	})
	commonChain := alice.New(c.Handler, middleware.RecoverMiddleware)

	a := middleware.NewAuthMiddleware(db, firebaseClient)
	authChain := commonChain.Append(a.Handler)

	ebookUsecase := usecase.NewEbookUseCase(db)
	ebookHandler := handler.NewEbookHandler(ebookUsecase)

	r := mux.NewRouter()
	sr := r.PathPrefix("/v1").Subrouter()
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/ping").Handler(commonChain.Then(AppHandler{ping}))
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/list-ebooks").Handler(authChain.Then(AppHandler{ebookHandler.Index}))

	return r
}

func ping(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, "pong", nil
}
