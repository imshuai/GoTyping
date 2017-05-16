package main

import (
	"errors"
	"time"
)

type Artical struct {
	ID         int64     `json:"id"`
	Name       string    `xorm:"varchar(255) index" json:"name"`
	Alias      string    `xorm:"varchar(255) index" json:"alias"`
	Summary    string    `xorm:"varchar(255) index" json:"summary"`
	Content    string    `xorm:"text" json:"content"`
	Author     string    `xorm:"varchar(30)" json:"author"`
	CreateTime time.Time `xorm:"created" json:"create_time"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"`
}

type ArticalSummary struct {
}

func NewArtical() *Artical {
	return &Artical{}
}

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

func (atl *Artical) Update() (bool, error) {
	return true, nil
}

func (atl *Artical) Delete() (bool, error) {
	return true, nil
}
