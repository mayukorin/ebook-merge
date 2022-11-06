package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/pkg/errors"
)

func InsertPurchase(db *sqlx.Tx, purchase *model.Purchase) (int64, error) {
	var id int64
	insertResult, err := db.Exec("insert into purchases (user_id, e_book_id) values (?, ?)", purchase.UserID, purchase.EBookID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert purchase")
	}
	id, err = insertResult.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "failed to select last insert id")
	}
	return id, nil
}

func FindPurchaseByUserIDAndEBookID(db *sqlx.DB, userId int64, eBookId int64) (*model.Purchase, error) {
	var purchase model.Purchase
	err := db.Get(&purchase, "select id, user_id, e_book_id from purchases where user_id = ? and e_book_id = ?", userId, eBookId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find ebook")
	}
	return &purchase, nil
}
