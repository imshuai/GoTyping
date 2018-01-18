package Model

import (
	"github.com/go-xorm/xorm"
)

var (
	db *xorm.Engine
)

func init() {
	db = nil
}
