package handler

import (
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/mayukorin/ebook-merge/app/interfaces/api/httputil"
	"github.com/mayukorin/ebook-merge/app/usecase"
	"github.com/mayukorin/ebook-merge/swagger/generated_swagger"
)

type GmailApiOAuth2TokenHandler struct {
	gmailApiOauth2TokenUseCase *usecase.GmailApiOauth2TokenUseCase
}

func NewGmailApiOAuth2TokenHandler(gmailApiOauth2TokenUseCase *usecase.GmailApiOauth2TokenUseCase) *GmailApiOAuth2TokenHandler {
	return &GmailApiOAuth2TokenHandler{gmailApiOauth2TokenUseCase: gmailApiOauth2TokenUseCase}
}

func (g *GmailApiOAuth2TokenHandler) GenerateGmailApiConsentPageURL(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	url, err := g.gmailApiOauth2TokenUseCase.GenerateGmailApiConsentPageURL()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, generated_swagger.GmailAPIConsentPageURL{GoogleConcentPageURL: &url}, nil
}

func (g *GmailApiOAuth2TokenHandler) CreateGmailApiOAuth2Token(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var req generated_swagger.CreateGmailAPIOAuth2TokenRequest
	if err := parseModelRequest(r, &req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	newOauth2TokenId, err := g.gmailApiOauth2TokenUseCase.GenerateAndInsertOAuth2Token(swag.StringValue(req.Code), user.ID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, generated_swagger.CreateGmailAPIOAuth2TokenResponse{ID: newOauth2TokenId}, nil
}
