package entity

type Entity interface {
	// テーブル名を取得する
	TableName() string
}
