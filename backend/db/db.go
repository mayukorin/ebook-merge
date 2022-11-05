package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDB(mysqlDSN string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", mysqlDSN)
	if err != nil {
		return nil, fmt.Errorf("failed open mysql. %s", err)
	}
	return db, nil
}
