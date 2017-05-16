package main

import (
	"time"
)

type Artical struct {
	ID         int64
	Name       string    `xorm:"varchar(255) index"`
	Summary    string    `xorm:"varchar(255) index"`
	Content    string    `xorm:"text"`
	Author     string    `xorm:"varchar(30)"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func NewArtical() *Artical {
	return &Artical{}
}

func (atl *Artical) Insert() *Artical {

}
