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
)

type Server struct {
	db             *sqlx.DB
	router         *mux.Router
	firebaseClient *auth.Client
}

func NewServer(mysqlDSN string, firebaseServiceAccountKeyPath string) (*Server, error) {
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

	s.router = NewRouter(s.db, s.firebaseClient)

	return s, nil
}

func (s *Server) Run(port string) {
	http.ListenAndServe(":"+port, s.router)
}
