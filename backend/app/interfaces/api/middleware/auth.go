package middleware

import (
	"fmt"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/repository"
	"github.com/mayukorin/ebook-merge/app/interfaces/api/httputil"
	"github.com/mayukorin/ebook-merge/firebase"
)

type AuthMiddleware struct {
	db             *sqlx.DB
	firebaseClient *auth.Client
}

func NewAuthMiddleware(db *sqlx.DB, firebaseClient *auth.Client) *AuthMiddleware {
	return &AuthMiddleware{
		db:             db,
		firebaseClient: firebaseClient,
	}
}

func (a *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idToken := getIdTokenFromHeader(r)
		token, err := a.firebaseClient.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			httputil.RespondErrByJSON(w, r, http.StatusForbidden, fmt.Errorf("unverified id token: %w", err))
			return
		}
		userData, err := a.firebaseClient.GetUser(r.Context(), token.UID)
		if err != nil {

			httputil.RespondErrByJSON(w, r, http.StatusForbidden, fmt.Errorf("failed to get user data from firebase: %w", err))
			return
		}

		firebaseUser := firebase.TransformToFirebaseUser(userData)
		err = repository.SyncUserWithFirebase(a.db, &firebaseUser)
		if err != nil {
			httputil.RespondErrByJSON(w, r, http.StatusForbidden, fmt.Errorf("failed to sync user: %w", err))
			return
		}
		user, err := repository.GetUser(a.db, token.UID)
		if err != nil {
			httputil.RespondErrByJSON(w, r, http.StatusForbidden, fmt.Errorf("failed to get user: %w", err))
			return
		}

		ctx := httputil.ContextWithUser(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getIdTokenFromHeader(r *http.Request) string {
	bearIdToken := r.Header.Get("Authorization")
	return bearIdToken[7:]
}
