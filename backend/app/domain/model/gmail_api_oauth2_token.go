package model

import (
	"time"

	"golang.org/x/oauth2"
)

type GmailApiOauth2Token struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	Email        string    `db:"email"`
	AccessToken  string    `db:"access_token"`
	TokenType    string    `db:"token_type"`
	RefreshToken string    `db:"refresh_token"`
	Expiry       time.Time `db:"expiry"`
}

func (g *GmailApiOauth2Token) ConvertToOauth2Token() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  g.AccessToken,
		TokenType:    g.TokenType,
		RefreshToken: g.RefreshToken,
		Expiry:       g.Expiry,
	}
}
