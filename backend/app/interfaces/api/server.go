package api

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/db"
)

type Server struct {
	db     *sqlx.DB
	router *mux.Router
}

func NewServer(mysqlDSN string) (*Server, error) {
	s := &Server{}

	db, err := db.NewDB(mysqlDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}
	s.db = db

	s.router = NewRouter()

	return s, nil
}

func (s *Server) Run(port string) {
	http.ListenAndServe(":"+port, s.router)
}
