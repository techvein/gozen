package oauth

import (
	"errors"

	"github.com/google/go-github/github"
	"github.com/google/logger"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"

	"gozen/config"
)

type OAuthGithub struct {
	*github.User
}

// Github OAuth 設定
var githubScopes = []string{
	"user:email",
	"repo",
}

var (
	// You must register the app at https://github.com/settings/applications
	// Set callback to http://127.0.0.1:7000/github_oauth_cb
	// Set ClientId and ClientSecret to
	oauthConf = &oauth2.Config{
		ClientID:     config.Oauth.Github.ClientID,
		ClientSecret: config.Oauth.Github.ClientSecret,
		Scopes:       githubScopes,
		Endpoint:     githuboauth.Endpoint,
	}
	// random string for oauth2 API calls to protect against CSRF
	oauthStateString = "thisshouldberandom"
)

func NewOAuthGitHub() User {
	return new(OAuthGithub)
}

// リダイレクトURLを作成する
func (self *OAuthGithub) GenerateLoginUrl() string {
	return oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
}

// CallBack処理を行う
func (self *OAuthGithub) Callback(state string, code string) (User, error) {
	if state != oauthStateString {
		logger.Errorf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		return nil, errors.New("invalid oauth state")
	}

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		logger.Errorf("oauthConf.Exchange() failed with '%s'\n", err)
		return nil, errors.New("oauthConf.Exchange() failed")
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")

	// Note: The returned email is the user's publicly visible email address (or null if the user has not specified a public email address in their profile
	// https://developer.github.com/v3/users/#get-a-single-user
	if user.Email == nil {
		// https://godoc.org/github.com/google/go-github/github#UsersService.ListEmails
		emails, _, _ := client.Users.ListEmails(nil)
		for _, email := range emails {
			logger.Infoln("Email:", *email.Email, "Primary:", *email.Primary, "Verified:", *email.Verified)
			user.Email = email.Email
			break
		}
	}

	if err != nil {
		logger.Errorf("client.Users.Get() faled with '%s'\n", err)
		return nil, errors.New("client.Users.Get() faled")
	}

	return &OAuthGithub{user}, nil
}

func (self *OAuthGithub) GetID() *int {
	return self.ID
}
func (self *OAuthGithub) GetName() *string {
	return self.Name
}
func (self *OAuthGithub) GetEmail() *string {
	return self.Email
}
func (self *OAuthGithub) GetSource() string {
	return "github"
}
func (self *OAuthGithub) GetClientID() *string {
	clientID := config.Oauth.Github.ClientID
	if clientID == "" {
		return nil
	}
	return &clientID
}
func (self *OAuthGithub) GetClientSecret() *string {
	clientSecret := config.Oauth.Github.ClientSecret
	if clientSecret == "" {
		return nil
	}
	return &clientSecret
}
