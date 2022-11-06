package api

import (
	"fmt"
	"net/http"

	"firebase.google.com/go/auth"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/db"
	"github.com/mayukorin/ebook-merge/firebase"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	db             *sqlx.DB
	router         *mux.Router
	firebaseClient *auth.Client
}

func NewServer(mysqlDSN string, firebaseServiceAccountKeyPath string, oauth2ClientID string, oauth2ClientSecret string, oauth2RedirectURL string) (*Server, error) {
	s := &Server{}

	db, err := db.NewDB(mysqlDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}
	s.db = db

	fc, err := firebase.NewFirebaseClient(firebaseServiceAccountKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to init firebase client: %w", err)
	}
	s.firebaseClient = fc

	oauth2Config := &oauth2.Config{
		ClientID:     oauth2ClientID,
		ClientSecret: oauth2ClientSecret,
		RedirectURL:  oauth2RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	s.router = NewRouter(s.db, s.firebaseClient, oauth2Config)

	return s, nil
}

func (s *Server) Run(port string) {
	http.ListenAndServe(":"+port, s.router)
}
