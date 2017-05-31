package main

//Option 定义站点配置信息
type Option struct {
	ID       int64
	SiteName string `xorm:"unique"`
}
