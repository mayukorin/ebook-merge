package firebase

import "firebase.google.com/go/auth"

type FirebaseUser struct {
	FirebaseUID string
	Email       string
}

func TransformToFirebaseUser(userRecord *auth.UserRecord) FirebaseUser {
	return FirebaseUser{
		FirebaseUID: userRecord.UID,
		Email:       userRecord.Email,
	}
}
