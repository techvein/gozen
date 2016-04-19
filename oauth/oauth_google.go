package oauth

import (
	ejson "encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"gozen/config"
)

type oAuthGoogle struct {
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
	return new(oAuthGoogle)
}

// リダイレクトURLを作成する
func (self *oAuthGoogle) GenerateLoginUrl() string {
	return conf.AuthCodeURL("state")
}

// CallBack処理を行う
func (self *oAuthGoogle) Callback(state string, code string) (User, error) {
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

func (self *oAuthGoogle) GetID() *int {
	id, _ := strconv.Atoi(self.Id)
	return &id
}
func (self *oAuthGoogle) GetName() *string {
	return &self.Name
}
func (self *oAuthGoogle) GetEmail() *string {
	return &self.Email
}
func (self *oAuthGoogle) GetSource() string {
	return "google"
}
func (self *oAuthGoogle) GetClientID() *string {
	clientID := config.Oauth.Google.ClientID
	if clientID == "" {
		return nil
	}
	return &clientID
}
func (self *oAuthGoogle) GetClientSecret() *string {
	clientSecret := config.Oauth.Google.ClientSecret
	if clientSecret == "" {
		return nil
	}
	return &clientSecret
}
