package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/logger"
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
		logger.Infoln("環境変数:" + gozenEnvKey + "が設定されていません。" + DevelopmentStr + "として動作します。")
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

func init() {
	gopath := os.Getenv("GOPATH")
	env := GetEnv()
	var err error

	// conf.<環境名>.ymlファイルから読み込む
	filePath := filepath.Join(gopath, "src/gozen/config/environment", "conf."+env.Name()+".yml")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Fatal(err)
	}

	var conf confYml
	err = yaml.Unmarshal([]byte(file), &conf)
	if err != nil {
		logger.Fatal(err)
	}
	Db = conf.Database
	Oauth = conf.Oauth

	// "file_path"の設定がなければgopath/log/gozen.logに設定する
	// "verbose"の設定がない場合はfalseになる
	if conf.Log.FilePath == "" {
		conf.Log.FilePath = "/var/log/gozen.log"
	}
	Log = conf.Log
}

type confYml struct {
	Mode     string      `yaml:"mode"`
	Database databaseYml `yaml:"database"`
	Oauth    oauthYml    `yaml:"oauth"`
	Log      logYml      `yaml:"log"`
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
	Github clientYml
	Google clientYml
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
