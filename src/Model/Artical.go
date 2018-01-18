package Model

import (
	"errors"
	"time"
)

//Artical 定义文章类
type Artical struct {
	ID         int64  `json:"id"`
	Name       string `xorm:"varchar(255) index" json:"name"`
	Alias      string `xorm:"varchar(255) index" json:"alias"`
	Summary    string `xorm:"varchar(1024) index" json:"summary"`
	Content    string `xorm:"text" json:"content"`
	Author     string `xorm:"varchar(30)" json:"author"`
	Category   int64
	Keywords   []string
	CreateTime time.Time `xorm:"created" json:"create_time"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"`
}

//Query 根据ID获取文章
func (atl *Artical) Query() (bool, error) {
	return db.ID(atl.ID).Get(atl)
}

//Insert 将新文章插入数据库
func (atl *Artical) Insert() (bool, error) {
	if atl.ID != 0 {
		return false, errors.New("artical already been inserted")
	}
	effect, err := db.InsertOne(atl)
	if err != nil || effect != 1 {
		return false, err
	}
	return true, nil
}

//Update 更新文章信息
func (atl *Artical) Update() (bool, error) {
	return true, nil
}

//Delete 删除指定文章
func (atl *Artical) Delete() (bool, error) {
	return true, nil
}

//Count 文章计数
func (atl *Artical) Count() int64 {
	count, err := db.OrderBy("id").Count(new(Artical))
	if err != nil {
		return 0
	}
	return count
}
