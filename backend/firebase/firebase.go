package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFirebaseClient(firebaseServiceAccountKeyPath string) (*auth.Client, error) {
	opt := option.WithCredentialsFile(firebaseServiceAccountKeyPath)
	fapp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to init firebase app: %w", err)
	}

	fclient, err := fapp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to init firebase client: %w", err)
	}

	return fclient, nil

}
