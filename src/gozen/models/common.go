package models

import (
	"github.com/gocraft/dbr"
)

type CommonModel struct {
	Id        dbr.NullInt64
	CreatedAt dbr.NullTime
	UpdatedAt dbr.NullTime
}
