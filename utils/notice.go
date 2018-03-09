package utils

import (
	"io/ioutil"

	"github.com/imshuai/serverchan"
)

//NoticManager 通过ServerChan通知管理员消息
func NoticManager(title, content string) {
	if sc != nil {
		sc.Send(title, content)
	}
}

//serverchanInit 初始化Serverchan通讯组件
func serverchanInit() {
	scSecretKey := func() string {
		byts, err := ioutil.ReadFile("serverchan.token")
		if err != nil {
			LogError("read serverchan.token fail with error:", err)
			return ""
		}
		return string(byts)
	}()
	if scSecretKey != "" {
		sc = serverchan.NewServerChan(scSecretKey)
	}
}
