package google_login

import (
	"burmese_jewellery/config"
	"burmese_jewellery/env"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	oauth2State            = "burmese_jewellery"
	oauth2UserInfoEndPoint = "https://www.googleapis.com/oauth2/v2/userinfo"
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     config.Get().Google.Oauth2Config.ClientID,
		ClientSecret: config.Get().Google.Oauth2Config.ClientSecret,
		RedirectURL: fmt.Sprintf(
			"%s://%s:%d/%s",
			env.Get().Http.Protocol,
			env.Get().Http.Domain,
			env.Get().Http.Port,
			config.Get().Google.Oauth2Config.RedirectURL,
		),
		Scopes:   []string{"profile", "email"},
		Endpoint: google.Endpoint,
	}
)

type UserInfo struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}

func GetLoginURL() string {
	return oauth2Config.AuthCodeURL(oauth2State)
}

func Callback(ctx context.Context, state string, code string) (*UserInfo, error) {

	if state != oauth2State {
		return nil, errors.New("oauth state does not match")
	}

	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	resp, err := oauth2Config.Client(ctx, token).Get(oauth2UserInfoEndPoint)
	if err != nil {
		return nil, err
	}

	var userInfo *UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return userInfo, err

}
