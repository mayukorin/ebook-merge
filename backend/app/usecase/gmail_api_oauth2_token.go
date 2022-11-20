package usecase

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/mayukorin/ebook-merge/app/domain/repository"
	"github.com/mayukorin/ebook-merge/app/interfaces/api/google"
	"github.com/mayukorin/ebook-merge/db/dbutil"
	"golang.org/x/oauth2"
)

type GmailApiOauth2TokenUseCase struct {
	db   *sqlx.DB
	conf *oauth2.Config
}

func NewGmailApiOauth2TokenUseCase(db *sqlx.DB, conf *oauth2.Config) *GmailApiOauth2TokenUseCase {
	return &GmailApiOauth2TokenUseCase{
		db:   db,
		conf: conf,
	}
}

func (g *GmailApiOauth2TokenUseCase) GenerateAndInsertOAuth2Token(code string, userId int64) (int64, error) {
	token, err := g.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return 0, fmt.Errorf("failed to generate oauth2 token from code : %w", err)
	}
	email, err := google.GetEmailOfOAuth2Token(token, g.conf)

	newOauth2Token := &model.GmailApiOauth2Token{
		UserID:       userId,
		Email:        email,
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	var createdId int64
	if err := dbutil.TXHandler(g.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertOAuth2Token(tx, newOauth2Token)
		if err != nil {
			return err
		}
		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, fmt.Errorf("failed oauth2_token insert transaction: %w", err)
	}
	return createdId, nil
}

func (g *GmailApiOauth2TokenUseCase) GenerateGmailApiConsentPageURL() (string, error) {
	// TODO: 0 → useId に変える
	authCodeUrlOptions := []oauth2.AuthCodeOption{oauth2.SetAuthURLParam("access_type", "offline"), oauth2.SetAuthURLParam("approval_prompt", "force")}
	url := g.conf.AuthCodeURL("0", authCodeUrlOptions...)
	return url, nil
}

// func (g *GmailApiOauth2TokenUseCase) GenerateOAuth2Token(code string) (*oauth2.Token, error) {

// 	token, err := g.conf.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate oauth2 token from code : %w", err)
// 	}
// 	return token, err
// }

// func (g *GmailApiOauth2TokenUseCase) GenerateOAuth2Client(token *oauth2.Token) *http.Client {
// 	client := g.conf.Client(oauth2.NoContext, token)
// 	return client
// }

// func (g *GmailApiOauth2TokenUseCase) GenerateRequestForOauth2UserInfo(token *oauth2.Token) *http.Request {
// 	reqForOAuth2UserInfo, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo", nil)
// 	reqForOAuth2UserInfo.Header.Set("Authorization", "Bearer "+token.AccessToken)
// 	return reqForOAuth2UserInfo
// }

// func (g *GmailApiOauth2TokenUseCase) Insert(userId int64, accessToken string, tokenType string, refreshToken string, expirty time.Time, email string) (int64, error) {
// 	newOauth2Token := &model.GmailApiOauth2Token{
// 		UserID:       userId,
// 		Email:        email,
// 		AccessToken:  accessToken,
// 		TokenType:    tokenType,
// 		RefreshToken: refreshToken,
// 		Expiry:       expirty,
// 	}

// 	var createdId int64
// 	if err := dbutil.TXHandler(g.db, func(tx *sqlx.Tx) error {
// 		id, err := repository.InsertOAuth2Token(tx, newOauth2Token)
// 		if err != nil {
// 			return err
// 		}
// 		createdId = id
// 		if err := tx.Commit(); err != nil {
// 			return err
// 		}
// 		return err
// 	}); err != nil {
// 		return 0, fmt.Errorf("failed oauth2_token insert transaction: %w", err)
// 	}
// 	return createdId, nil
// }
