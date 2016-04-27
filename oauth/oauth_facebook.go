package oauth

import (
	ejson "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"

	"gozen/config"
	"gozen/models/json"
)

type OAuthFacebook struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
	Gender    string `json:"gender"`
	Locale    string `json:"locale"`
	conf      *oauth2.Config
}

type facebookAccessToken struct {
	Error       *json.ErrorJSON `json:"error"`
	AccessToken *string         `json:"access_token"`
}

// https://developers.facebook.com/docs/graph-api/reference/v2.6/debug_token/
type facebookData struct {
	Data facebookDataContent `json:"data"`
}
type facebookDataContent struct {
	AppID       string `json:"app_id"`
	Application string `json:"application"`
	IsValid     string `json:"is_valid"`
	UserID      string `json:"user_id"`
}

// Facebook OAuth 設定
var facebookScopes = []string{
	"email",
}

func NewOAuthFacebook() User {
	fb := new(OAuthFacebook)
	fb.conf = &oauth2.Config{
		ClientID:     config.Oauth.Facebook.ClientID,
		ClientSecret: config.Oauth.Facebook.ClientSecret,
		RedirectURL:  config.Oauth.Facebook.RedirectURL,
		Scopes:       facebookScopes,
		Endpoint:     facebook.Endpoint,
	}
	return fb
}

// リダイレクトURLを作成する
func (self *OAuthFacebook) GenerateLoginUrl() string {
	return self.conf.AuthCodeURL("")
}

// CallBack処理を行う
func (self *OAuthFacebook) Callback(state string, code string) (User, error) {
	tok, err := self.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	client := conf.Client(oauth2.NoContext, tok)

	// get access_token
	accessToken := &facebookAccessToken{}
	err = self.request(client, accessToken,
		fmt.Sprintf("/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials",
			self.conf.ClientID, self.conf.ClientSecret))
	if err != nil {
		return nil, err
	}
	if accessToken.Error != nil {
		return nil, errors.New(accessToken.Error.Message)
	}

	// https://developers.facebook.com/docs/graph-api/reference/v2.6/debug_token/
	facebookToken := "" // TODO
	if _, err := self.CheckToken(client, accessToken, facebookToken); err != nil {
		return nil, err
	}

	return self, nil
}

func (self *OAuthFacebook) CheckToken(client *http.Client, accessToken *facebookAccessToken, facebookToken string) (bool, error) {
	facebookData := &facebookData{}
	err := self.request(client, facebookData,
		fmt.Sprintf(
			"/debug_token?input_token=%s&access_token=%s",
			facebookToken,
			*accessToken.AccessToken,
		))
	if err != err {
		return false, errors.New("request error.")
	}

	logger.Infoln(facebookData.Data.AppID)
	if facebookData.Data.AppID != self.conf.ClientID {
		return false, errors.New("The facebookToken is wrong.")
	}
	return true, nil
}

func (self *OAuthFacebook) request(client *http.Client, result interface{}, path string) error {
	response, err := client.Get(fmt.Sprintf("https://graph.facebook.com/v2.6%s", path))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	logger.Infoln(string(contents))
	if err != nil {
		return err
	}
	ejson.Unmarshal(contents, result)
	return nil
}

func (self *OAuthFacebook) GetID() *int {
	id, _ := strconv.Atoi(self.Id)
	return &id
}
func (self *OAuthFacebook) GetName() *string {
	return &self.Name
}
func (self *OAuthFacebook) GetEmail() *string {
	return &self.Email
}
func (self *OAuthFacebook) GetSource() string {
	return "facebook"
}
func (self *OAuthFacebook) GetClientID() *string {
	clientID := config.Oauth.Facebook.ClientID
	if clientID == "" {
		return nil
	}
	return &clientID
}
func (self *OAuthFacebook) GetClientSecret() *string {
	clientSecret := config.Oauth.Facebook.ClientSecret
	if clientSecret == "" {
		return nil
	}
	return &clientSecret
}
