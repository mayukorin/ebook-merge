package repository

import (
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/mayukorin/ebook-merge/firebase"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func GetUser(db *sqlx.DB, uid string) (*model.User, error) {
	var u model.User
	if err := db.Get(&u, `
select id, firebase_uid, email from users where firebase_uid = ?
limit 1
	`, uid); err != nil {
		return nil, errors.Wrap(err, "failed to execute get user")
	}
	return &u, nil
}

func SyncUserWithFirebase(db *sqlx.DB, fu *firebase.FirebaseUser) error {
	_, err := db.Exec(`
insert into users (firebase_uid, email)
values (?, ?)
on DUPLICATE KEY 
	update email = ?
`, fu.FirebaseUID, fu.Email, fu.Email)
	if err != nil {
		return errors.Wrap(err, "failed to execute sync user with firebase")
	}
	return nil
}
