package oauth

import (
	ejson "encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/techvein/gozen/config"
)

type OAuthGoogle struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail string `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`
	Locale        string `json:"locale"`
}

// Google OAuth 設定
var googleScopes = []string{
	"profile",
	"email",
}

var conf = &oauth2.Config{
	ClientID:     config.Oauth.Google.ClientID,
	ClientSecret: config.Oauth.Google.ClientSecret,
	RedirectURL:  config.Oauth.Google.RedirectURL,
	Scopes:       googleScopes,
	Endpoint:     google.Endpoint,
}

func NewOAuthGoogle() User {
	return new(OAuthGoogle)
}

// リダイレクトURLを作成する
func (self *OAuthGoogle) GenerateLoginUrl() string {
	return conf.AuthCodeURL("state")
}

// CallBack処理を行う
func (self *OAuthGoogle) Callback(state string, code string) (User, error) {
	var tok, err = conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("%v", err)
	}
	var client = conf.Client(oauth2.NoContext, tok)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	ejson.Unmarshal(contents, &self)

	return self, err
}

func (self *OAuthGoogle) GetID() *int {
	id, _ := strconv.Atoi(self.Id)
	return &id
}
func (self *OAuthGoogle) GetName() *string {
	return &self.Name
}
func (self *OAuthGoogle) GetEmail() *string {
	return &self.Email
}
func (self *OAuthGoogle) GetSource() string {
	return "google"
}
func (self *OAuthGoogle) GetClientID() *string {
	clientID := config.Oauth.Google.ClientID
	if clientID == "" {
		return nil
	}
	return &clientID
}
func (self *OAuthGoogle) GetClientSecret() *string {
	clientSecret := config.Oauth.Google.ClientSecret
	if clientSecret == "" {
		return nil
	}
	return &clientSecret
}
