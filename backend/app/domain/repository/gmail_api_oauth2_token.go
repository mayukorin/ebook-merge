package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/pkg/errors"
)

func InsertOAuth2Token(db *sqlx.Tx, oauth2Token *model.GmailApiOauth2Token) (int64, error) {
	stmt, err := db.Preparex("insert into gmail_api_oauth2_tokens (user_id, email, access_token, token_type, refresh_token, expiry) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, errors.Wrap(err, "failed to set perpared statement")
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	insertResult, err := stmt.Exec(oauth2Token.UserID, oauth2Token.Email, oauth2Token.AccessToken, oauth2Token.TokenType, oauth2Token.RefreshToken, oauth2Token.Expiry)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert oauth2_token")
	}
	id, err = insertResult.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "failed to select last insert id")
	}
	return id, nil
}

func SelectGmailApiOAuth2TokenByUserID(db *sqlx.DB, userId int64) ([]model.GmailApiOauth2Token, error) {
	var oauth2Tokens []model.GmailApiOauth2Token
	err := db.Select(&oauth2Tokens, "select id, email, access_token, token_type, refresh_token, expiry from gmail_api_oauth2_tokens where user_id = ?", userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select gmail_api_oauth2_tokens by user id")
	}
	return oauth2Tokens, nil
}
