package models

import (
	"github.com/gocraft/dbr"
)

type Auth struct {
	Id        uint64 `db:"id"`
	UserId    uint64 `db:"user_id"`
	Source    string `db:"source"`
	SourceId  int    `db:"source_id"`
	Email     string `db:"email"`
	CreatedAt dbr.NullTime
	UpdatedAt dbr.NullTime
}

func (self *Auth) TableName() string {
	return "auths"
}

func NewAuth() *Auth {
	return new(Auth)
}
