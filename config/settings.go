package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const gozenEnvKey string = "GOZEN_ENV"

const (
	ProductionStr  string = "production"
	StagingStr     string = "staging"
	DevelopmentStr string = "development"
)

type Env uint

const (
	Production Env = 1 << iota
	Staging
	Development
)

var gozenEnv Env = 0

func GetEnv() Env {
	if gozenEnv != 0 {
		return gozenEnv
	}

	env := os.Getenv(gozenEnvKey)
	if env == "" {
		log.Println("環境変数:" + gozenEnvKey + "が設定されていません。" + DevelopmentStr + "として動作します。")
		gozenEnv = Development
		return gozenEnv
	}

	switch env {
	case ProductionStr:
		gozenEnv = Production
	case StagingStr:
		gozenEnv = Staging
	case DevelopmentStr:
		gozenEnv = Development
	}
	return gozenEnv
}

func (env Env) Name() string {
	switch env {
	case Production:
		return ProductionStr
	case Staging:
		return StagingStr
	case Development:
		return DevelopmentStr
	}
	return "環境が不明です"
}

var Db databaseYml
var Oauth oauthYml
var Log logYml
var Push pushYml

func init() {
	gopath := os.Getenv("GOPATH")
	env := GetEnv()
	var err error

	// conf.<環境名>.ymlファイルから読み込む
	filePath := filepath.Join(gopath, "src/github.com/techvein/gozen/config/environment", "conf."+env.Name()+".yml")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var conf confYml
	err = yaml.Unmarshal([]byte(file), &conf)
	if err != nil {
		log.Fatal(err)
	}
	Db = conf.Database
	Oauth = conf.Oauth
	Push = conf.Push

	// "file_path"の設定がなければgopath/log/gozen.logに設定する
	// "verbose"の設定がない場合はfalseになる
	if conf.Log.FilePath == "" {
		ex, err := os.Executable()
	    if err != nil {
	        panic(err)
	    }
	    exPath := path.Dir(ex)
	    fmt.Println(exPath)
		conf.Log.FilePath =  exPath + "/gozen.log"
	}
	Log = conf.Log
}

type confYml struct {
	Mode     string      `yaml:"mode"`
	Database databaseYml `yaml:"database"`
	Oauth    oauthYml    `yaml:"oauth"`
	Log      logYml      `yaml:"log"`
	Push     pushYml     `yaml:"push"`
}

type databaseYml struct {
	Adapter  string
	Database string
	Username string
	Password string
	Host     string
	Port     string
}

type oauthYml struct {
	// oauth認証後に遷移するurl
	AfterOauthUrl string `yaml:"after_oauth_url"`
	Github        clientYml
	Google        clientYml
	Facebook      clientYml
}

type clientYml struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
	Scopes       []string
}

type logYml struct {
	Verbose  bool   `yaml:"verbose"`
	FilePath string `yaml:"file_path"`
}
type pushYml struct {
	Gcm gcmYml `yaml:"gcm"`
}

type gcmYml struct {
	ApiKey string `yaml:"api_key"`
}
