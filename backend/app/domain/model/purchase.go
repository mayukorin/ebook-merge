package model

type Purchase struct {
	ID      int64 `db:"id"`
	UserID  int64 `db:"user_id"`
	EBookID int64 `db:"e_book_id"`
}
