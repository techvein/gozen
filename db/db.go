package db

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"

	"github.com/techvein/gozen/config"
)

var session *dbr.Session
var once = new(sync.Once)

func GetSession() *dbr.Session {
	// 最初の一度だけ呼ぶ
	once.Do(func() {
		session = connect()
	})
	return session
}

func connect() *dbr.Session {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Db.Username,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Database)

	conn, err := dbr.Open(config.Db.Adapter, dsn+"?charset=utf8&parseTime=True", nil)
	if err != nil {
		log.Fatalln(err)
	}

	return conn.NewSession(nil)
}
