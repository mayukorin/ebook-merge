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
	"golang.org/x/oauth2"
)

func NewRouter(db *sqlx.DB, firebaseClient *auth.Client, gmailOauth2Config *oauth2.Config) *mux.Router {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "Authorization"},
	})
	commonChain := alice.New(c.Handler, middleware.RecoverMiddleware)

	a := middleware.NewAuthMiddleware(db, firebaseClient)
	authChain := commonChain.Append(a.Handler)

	gmailApiOauth2TokenUsecase := usecase.NewGmailApiOauth2TokenUseCase(db, gmailOauth2Config)
	gmailApiOauth2TokenHandler := handler.NewGmailApiOAuth2TokenHandler(gmailApiOauth2TokenUsecase)

	ebookUsecase := usecase.NewEbookUseCase(db, gmailOauth2Config)
	ebookHandler := handler.NewEbookHandler(ebookUsecase)

	r := mux.NewRouter()
	sr := r.PathPrefix("/v1").Subrouter()
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/ping").Handler(commonChain.Then(AppHandler{ping}))
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/list-ebooks").Handler(authChain.Then(AppHandler{ebookHandler.Index}))
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/generate-consent-page-url-of-gmail-api").Handler(commonChain.Then(AppHandler{gmailApiOauth2TokenHandler.GenerateGmailApiConsentPageURL}))
	sr.Methods(http.MethodPost, http.MethodOptions).Path("/generate-oauth2-token-of-gmail-api").Handler(authChain.Then(AppHandler{gmailApiOauth2TokenHandler.CreateGmailApiOAuth2Token}))
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/scan-ebooks").Handler(authChain.Then(AppHandler{ebookHandler.ScanAllEbooksFromGmail}))
	sr.Methods(http.MethodGet, http.MethodOptions).Path("/test-scan-ebooks").Handler(commonChain.Then(AppHandler{ebookHandler.ScanTestEbooksFromGmail}))

	return r
}

func ping(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, "pong", nil
}
