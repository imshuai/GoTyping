package main

import (
	"errors"
	"time"
)

//Artical 定义文章类
type Artical struct {
	ID         int64  `json:"id"`
	Name       string `xorm:"varchar(255) index" json:"name"`
	Alias      string `xorm:"varchar(255) index" json:"alias"`
	Summary    string `xorm:"varchar(255) index" json:"summary"`
	Content    string `xorm:"text" json:"content"`
	Author     string `xorm:"varchar(30)" json:"author"`
	Category   int64
	Keywords   []string
	CreateTime time.Time `xorm:"created" json:"create_time"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"`
}

//NewArtical 创建新文章
func NewArtical() *Artical {
	return &Artical{}
}

//GetArtical 根据ID获取文章
func GetArtical(ID int64) *Artical {
	if ID < 0 {
		return &Artical{}
	} else if ID == 0 {
		return NewArtical()
	}
	atl := new(Artical)
	exist, err := db.ID(ID).Get(atl)
	if err != nil || !exist {
		return &Artical{}
	}
	return atl
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

func articalCount() int64 {
	count, err := db.OrderBy("id").Count(new(Artical))
	if err != nil {
		return 0
	}
	return count
}

//ArticalSummary 定义文章摘要类
type ArticalSummary struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Summary    string   `json:"summary"`
	Category   Category `json:"category"`
	CreateTime string   `json:"create_time"`
}

//GenSummary 生成文章摘要索引
func (atl *Artical) GenSummary() ArticalSummary {
	return ArticalSummary{
		ID:      atl.ID,
		Name:    atl.Name,
		Summary: atl.Summary,
		Category: func() Category {
			cat := new(Category)
			db.ID(atl.Category).Get(cat)
			return *cat
		}(),
		CreateTime: atl.CreateTime.Format("Monday, 02 Jan 2006"),
	}
}

//GetSummariesWithPageID 根据文章索引页获得摘要
func GetSummariesWithPageID(pid int) []ArticalSummary {
	count := articalCount()
	start := 0
	if start = (pid - 1) * 12; start < 0 || int64(start) > count {
		start = 0
	}
	atls := make([]*Artical, 0)
	db.Asc("id").Limit(12, start).Find(&atls)
	atlsums := make([]ArticalSummary, 0)
	for _, v := range atls {
		atlsums = append(atlsums, v.GenSummary())
	}
	return atlsums
}
