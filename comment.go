package main

import (
	"time"
)

type Comment struct {
	ID         int64 `json:"id"`
	ArticalID  int64
	Author     string    `xorm:"varchar(30)" json:"author"`
	Content    string    `xorm:"text" json:"content"`
	CreateTime time.Time `xorm:"created" json:"time"`
	CommentTo  int64     `xorm:"default(0)" json:"comment_to"`
	Like       int64     `xorm:"default(0)" json:"like"`
	Unlike     int64     `xorm:"default(0)" json:"unlike"`
}

type ArticalComments struct {
	ArticalID int64     `json:"artical_id"`
	Comments  []Comment `json:"comments"`
}

func NewComment() *Comment {
	return &Comment{}
}

func (cm *Comment) Insert() (bool, error) {
	if cm.ID != 0 {
		return false, errors.New("comment already been inserted")
	}
	effect, err := db.InsertOne(cm)
	if err != nil || effect != 1 {
		return false, err
	}
	return true, nil
}

func (cm *Comment) Delete() (bool, error) {
	return true, nil
}

func GetComments(articalID int64) (ArticalComments, error) {
	cms := make([]Comment, 0)
	err := db.Where("`artical_id`=?", articalID).Find(&cms)
	if err != nil {
		return ArticalComments{}, err
	}
	return ArticalComments{
		ArticalID: articalID,
		Comments:  cms,
	}, nil
}
