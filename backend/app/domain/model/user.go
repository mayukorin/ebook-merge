package model

type User struct {
	ID          int64  `db:"id"`
	FirebaseUID string `db:"firebase_uid"`
	Email       string `db:"email"`
}
