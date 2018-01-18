package Model

import (
	"errors"
	"time"
)

//Comment 定义评论类
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

//ArticalComments 定义文章评论列表类
type ArticalComments struct {
	ArticalID int64     `json:"artical_id"`
	Comments  []Comment `json:"comments"`
}

//NewComment 创建新评论
func NewComment() *Comment {
	return &Comment{}
}

//Insert 将新评论插入数据库
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

//Delete 删除指定评论
func (cm *Comment) Delete() (bool, error) {
	return true, nil
}

//GetComments 根据文章ID获取评论列表
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
