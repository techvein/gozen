package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/techvein/gozen/entity"
)

const tagName string = "db"

// structタグ("db")の付いた全部のタグ(カラム名)を取得する
func StructTagToColumns(entity entity.Entity) []string {

	var tableName string
	var rt reflect.Type
	switch f := entity.(type) {
	// TODO: case *User,*xxx・・・みたいにまとめたいが、まとめるとコンパイルエラーになる。
	case *User:
		tableName = f.TableName()
		rt = reflect.TypeOf(*f)
	default:

	}

	columns := make([]string, 0)

	for i, _ := range make([]int, rt.NumField()) {
		field := rt.Field(i)
		db := field.Tag.Get(tagName)

		if len(db) == 0 {
			continue
		}

		columns = append(columns, fmt.Sprintf("%s.%s", tableName, db))

	}

	return columns
}

// 現在時刻がtを過ぎているか
func IsNowAfter(t time.Time) bool {
	return time.Now().After(t)

}
