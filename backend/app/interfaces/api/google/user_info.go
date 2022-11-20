package google

import (
	"fmt"
	"net/http"

	"github.com/mayukorin/ebook-merge/swagger/generated_swagger"
	"golang.org/x/oauth2"
)

func GetEmailOfOAuth2Token(token *oauth2.Token, conf *oauth2.Config) (string, error) {
	client := conf.Client(oauth2.NoContext, token)

	reqForOAuth2UserInfo, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo", nil)
	if err != nil {
		return "", fmt.Errorf("failed to new request : %w", err)
	}
	reqForOAuth2UserInfo.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := client.Do(reqForOAuth2UserInfo)
	if err != nil {
		return "", fmt.Errorf("failed to fetch userinfo : %w", err)
	}
	var userInfo generated_swagger.OAuth2UserInfo
	if err = parseModelResponse(resp, &userInfo); err != nil {
		return "", fmt.Errorf("failed to parse userinfo : %w", err)
	}

	return userInfo.Email, nil
}
