package main

import (
	"golearn/lsyproject/liteblog/models"
	"encoding/gob"
	_ "golearn/lsyproject/liteblog/routers"
	_ "golearn/lsyproject/liteblog/models"
	"strings"
	"github.com/astaxie/beego"
)

func main() {
	initTemplate()
	initSession()
	beego.Run()
}

//自定义函数
func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
}

func initSession(){
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(&models.User{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}