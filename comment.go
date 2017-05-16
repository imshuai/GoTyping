package main

import (
	"time"
)

type Comment struct {
	ID         int64
	ArticalID  int64
	Author     string    `xorm:"varchar(30)"`
	Content    string    `xorm:"text"`
	CreateTime time.Time `xorm:"created"`
	CommentTo  int64     `xorm:"default(0)"`
	Like       int64     `xorm:"default(0)"`
	Unlike     int64     `xorm:"default(0)"`
}

func NewComment() *Comment {
	return &Comment{}
}
