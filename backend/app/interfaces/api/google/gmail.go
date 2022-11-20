package google

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailApi struct {
	service *gmail.Service
	ctx     context.Context
}

func NewGmailApi(token *oauth2.Token, conf *oauth2.Config) (*GmailApi, error) {
	gmailApi := GmailApi{}
	gmailApi.ctx = context.Background()
	client := conf.Client(gmailApi.ctx, token)
	srv, err := gmail.NewService(gmailApi.ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to new service:%w", err)
	}
	gmailApi.service = srv
	return &gmailApi, nil
}

func (g *GmailApi) SearhGmailMessages(searchWord string, email string) ([]*gmail.Message, error) {
	r, err := g.service.Users.Messages.List(email).Q(searchWord).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to search messages: %w", err)
	}
	return r.Messages, nil
}

func (g *GmailApi) GetGmailMessage(messageId string, email string) (*gmail.Message, error) {
	message, err := g.service.Users.Messages.Get(email, messageId).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get messages: %w", err)
	}
	return message, nil
}
