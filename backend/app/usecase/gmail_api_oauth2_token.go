package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/mayukorin/ebook-merge/app/domain/repository"
	"github.com/mayukorin/ebook-merge/app/domain/service"
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

func (g *GmailApiOauth2TokenUseCase) GenerateGmailApiConsentPageURL() (string, error) {
	// TODO: 0 → useId に変える
	authCodeUrlOptions := []oauth2.AuthCodeOption{oauth2.SetAuthURLParam("access_type", "offline")}
	url := g.conf.AuthCodeURL("0", authCodeUrlOptions...)
	return url, nil
}

func (g *GmailApiOauth2TokenUseCase) GenerateOAuth2Token(code string) (*oauth2.Token, error) {

	token, err := g.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("failed to generate oauth2 token from code : %w", err)
	}
	return token, err
}

func (g *GmailApiOauth2TokenUseCase) GenerateOAuth2Client(token *oauth2.Token) *http.Client {
	client := g.conf.Client(oauth2.NoContext, token)
	return client
}

func (g *GmailApiOauth2TokenUseCase) GenerateRequestForOauth2UserInfo(token *oauth2.Token) *http.Request {
	reqForOAuth2UserInfo, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo", nil)
	reqForOAuth2UserInfo.Header.Set("Authorization", "Bearer "+token.AccessToken)
	return reqForOAuth2UserInfo
}

func (g *GmailApiOauth2TokenUseCase) Insert(userId int64, accessToken string, tokenType string, refreshToken string, expirty time.Time, email string) (int64, error) {
	newOauth2Token := &model.GmailApiOauth2Token{
		UserID:       userId,
		Email:        email,
		AccessToken:  accessToken,
		TokenType:    tokenType,
		RefreshToken: refreshToken,
		Expiry:       expirty,
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

func (g *GmailApiOauth2TokenUseCase) ScanAllEbooksFromGmail(userId int64) error {
	gmailAPiOauth2Tokens, err := repository.SelectGmailApiOAuth2TokenByUserID(g.db, userId)
	if err != nil {
		return fmt.Errorf("failed gmail_api_oauth2_tokens by user id:%w", err)
	}

	kindleEbookService, err := repository.FindEbookServiceByName(g.db, "Kindle")
	if err != nil {
		return fmt.Errorf("failed kindle ebook service:%w", err)
	}

	bookLiveEbookService, err := repository.FindEbookServiceByName(g.db, "BookLive")
	if err != nil {
		return fmt.Errorf("failed BookLive ebook service:%w", err)
	}

	for _, gmailOauth2Token := range gmailAPiOauth2Tokens {
		bookLiveEbookTitles, err := service.ScanBookLiveEbooks(gmailOauth2Token, g.conf)
		if err != nil {
			return fmt.Errorf("failed scan ebook of booklive:%w", err)
		}

		for _, bookLiveTitle := range bookLiveEbookTitles {
			eBookId, err := service.FindOrCreateEbook(bookLiveTitle, bookLiveEbookService.ID, g.db)
			if err != nil {
				return fmt.Errorf("failed find or create ebook of book live:%w", err)
			}
			_, err = service.FindOrCreatePurchase(userId, eBookId, g.db)
			if err != nil {
				return fmt.Errorf("failed find or create purchase:%w", err)
			}
		}

		kindleEbookTitles, err := service.ScanKindleEbooks(gmailOauth2Token, g.conf)
		if err != nil {
			return fmt.Errorf("failed scan ebook of kindle:%w", err)
		}

		for _, kindleEbookTitle := range kindleEbookTitles {
			eBookId, err := service.FindOrCreateEbook(kindleEbookTitle, kindleEbookService.ID, g.db)
			if err != nil {
				return fmt.Errorf("failed find or create ebook of kindle:%w", err)
			}
			_, err = service.FindOrCreatePurchase(userId, eBookId, g.db)
			if err != nil {
				return fmt.Errorf("failed find or create purchase:%w", err)
			}
		}

	}
	return nil
}
