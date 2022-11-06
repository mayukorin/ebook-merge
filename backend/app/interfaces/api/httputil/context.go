package httputil

import (
	"context"
	"errors"

	"github.com/mayukorin/ebook-merge/app/domain/model"
)

func ContextWithUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, "user", user)
}

func GetUserFromContext(ctx context.Context) (*model.User, error) {
	v := ctx.Value("user")
	user, ok := v.(*model.User)
	if !ok {
		return nil, errors.New("failed to find user")
	}
	return user, nil
}
