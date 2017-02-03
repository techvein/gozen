package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"bitbucket.org/liamstask/goose/lib/goose"
	"gopkg.in/yaml.v2"

	"gozen/config"
)

var gopath string

func main() {
	gopath = os.Getenv("GOPATH")

	quitInstalll := make(chan bool)
	quitRunMigration := make(chan bool)

	go installLibraries(quitInstalll)

	go runMigration(quitRunMigration)

	<-quitInstalll
	<-quitRunMigration
}

func installLibraries(quit chan bool) {
	os.Chdir(filepath.Join(gopath, "src/gozen"))

	_, err := exec.LookPath("dep")
	if err != nil {
		log.Fatal("Please install govendor that is a package management tool.")
	}

	cmd := exec.Command("dep", "ensure", "-update", "-v")
	fmt.Print("run", cmd.Path, "ensure -update -v ...")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done dep ensure update.")
	quit <- true
}

func runMigration(quit chan bool) {
	runMigrationWithGoose(quit)
}

func runMigrationWithGoose(quit chan bool) {
	fmt.Println("run goose up")
	// config/environment から db/dbconf.ymlを作成
	dbconfYml := map[string]interface{}{
		config.GetEnv().Name(): map[interface{}]interface{}{
			"driver": config.Db.Adapter,
			"open": fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
				config.Db.Username,
				config.Db.Password,
				config.Db.Host,
				config.Db.Port,
				config.Db.Database,
			),
		},
	}

	data, err := yaml.Marshal(dbconfYml)
	if err != nil {
		log.Fatal(err)
	}

	p := filepath.Join(gopath, "src/gozen/db")

	ioutil.WriteFile(filepath.Join(p, "dbconf.yml"), data, os.ModePerm)

	// migration with goose
	conf, err := goose.NewDBConf(p, config.GetEnv().Name(), "")

	// goose up
	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	// goose down
	//current, err := goose.GetDBVersion(conf)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//target, err := goose.GetPreviousDBVersion(conf.MigrationsDir, current)
	//if err != nil {
	//	log.Fatal(err)
	//}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done goose up.")
	quit <- true
}
