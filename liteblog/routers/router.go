package routers

import (
	"golearn/lsyproject/liteblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
	)	
}
