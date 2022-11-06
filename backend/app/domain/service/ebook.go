package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/mayukorin/ebook-merge/app/domain/repository"
	"github.com/mayukorin/ebook-merge/db/dbutil"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func ScanKindleEbooks(gmailOauth2Token model.GmailApiOauth2Token, conf *oauth2.Config) ([]string, error) {
	ctx := context.Background()
	oauth2Token := gmailOauth2Token.ConvertToOauth2Token()
	client := conf.Client(context.Background(), oauth2Token)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmailclient: %w", err)
	}

	r, err := srv.Users.Messages.List(gmailOauth2Token.Email).Q("subject:Amazon.co.jpでのご注文").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to list kindle messages: %w", err)
	}

	var ebookTitles []string
	for _, message := range r.Messages {
		messageWithPayload, err := srv.Users.Messages.Get(gmailOauth2Token.Email, message.Id).Do()
		if err != nil {
			return nil, fmt.Errorf("unable to get messageWithPayload: %w", err)
		}
		message_content := strings.Replace(messageWithPayload.Payload.Parts[0].Body.Data, "-", "+", -1)
		message_content = strings.Replace(message_content, "_", "/", -1)
		decoded_message_content, err := base64.StdEncoding.DecodeString(message_content)
		if err != nil {
			return nil, fmt.Errorf("unable to decode base64: %w", err)
		}

		re := regexp.MustCompile(`>\s([^\r]+)\r\n*>\sKindle版`)
		matches := re.FindAllStringSubmatch(string(decoded_message_content), -1)
		for _, match := range matches {
			ebookTitles = append(ebookTitles, match[1])
		}
	}
	return ebookTitles, nil
}

func ScanBookLiveEbooks(gmailOauth2Token model.GmailApiOauth2Token, conf *oauth2.Config) ([]string, error) {
	ctx := context.Background()
	oauth2Token := gmailOauth2Token.ConvertToOauth2Token()
	client := conf.Client(context.Background(), oauth2Token)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmailclient: %w", err)
	}
	r, err := srv.Users.Messages.List(gmailOauth2Token.Email).Q("subject:【ブックライブ】ご購入").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to list booklive messages: %w", err)
	}

	var ebookTitles []string
	for _, message := range r.Messages {
		messageWithPayload, err := srv.Users.Messages.Get(gmailOauth2Token.Email, message.Id).Do()
		if err != nil {
			return nil, fmt.Errorf("unable to get messagewithPayload: %w", err)
		}
		message_content := strings.Replace(messageWithPayload.Payload.Parts[0].Body.Data, "-", "+", -1)
		message_content = strings.Replace(message_content, "_", "/", -1)
		decoded_message_content, err := base64.StdEncoding.DecodeString(message_content)
		if err != nil {
			return nil, fmt.Errorf("unable to decode base64: %w", err)
		}

		re := regexp.MustCompile(`1\.\s([^\r]+)`)
		matches := re.FindAllStringSubmatch(string(decoded_message_content), -1)
		for _, match := range matches {
			ebookTitles = append(ebookTitles, match[1])
		}
	}
	return ebookTitles, nil
}

func FindOrCreateEbook(eBookTitle string, eBookServiceId int64, db *sqlx.DB) (int64, error) {

	eBook, err := repository.FindEbookByEbookServiceIDAndTitle(db, eBookServiceId, eBookTitle)
	if err != nil {
		if errors.Unwrap(errors.Unwrap(err)) != sql.ErrNoRows {
			return 0, fmt.Errorf("failed find ebook:%w", err)
		}
	}
	if eBook != nil {
		return eBook.ID, nil
	}

	newEbook := &model.Ebook{
		Title: eBookTitle,
		EbookService: model.EbookService{
			ID: eBookServiceId,
		},
	}

	var createdId int64
	if err := dbutil.TXHandler(db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertEbook(tx, newEbook)
		if err != nil {
			return err
		}
		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, fmt.Errorf("failed ebook insert transaction: %w", err)
	}
	return createdId, nil
}

func FindOrCreatePurchase(userId int64, eBookId int64, db *sqlx.DB) (int64, error) {

	purchase, err := repository.FindPurchaseByUserIDAndEBookID(db, userId, eBookId)
	if err != nil {
		if errors.Unwrap(errors.Unwrap(err)) != sql.ErrNoRows {
			return 0, fmt.Errorf("failed find ebook:%w", err)
		}
	}
	if purchase != nil {
		return purchase.ID, nil
	}

	newPurchase := &model.Purchase{
		UserID:  userId,
		EBookID: eBookId,
	}

	var createdId int64
	if err := dbutil.TXHandler(db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertPurchase(tx, newPurchase)
		if err != nil {
			return err
		}
		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, fmt.Errorf("failed purchase insert transaction: %w", err)
	}
	return createdId, nil
}
