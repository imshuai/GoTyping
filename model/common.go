package model

import "time"

type Common struct {
	CreateTime time.Time `xorm:"created" json:"create_time"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"`
}
