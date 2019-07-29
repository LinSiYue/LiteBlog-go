package controllers

import (
	"errors"
	"golearn/lsyproject/liteblog/syserror"
	"golearn/lsyproject/liteblog/models"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type NestPreparer interface{
	NextPrepare()
}

type BaseController struct {
	beego.Controller
	IsLogin bool //标识 用户是否登陆
	User *models.User //登陆的用户
}	

//识别请求的路径，然后获取路径
func (ctx *BaseController) Prepare()  {
	
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	//判断用户是否登陆，判断session中是否存在用户，存在就表示已登陆，不存在表示未登录
	ctx.IsLogin = false
	u, ok := ctx.GetSession(SESSION_USER_KEY).(*models.User)

	ctx.IsLogin = false
	if ok{
		ctx.User = u
		ctx.Data["User"] = ctx.User
		ctx.IsLogin = true
	}

	ctx.Data["IsLogin"] = ctx.IsLogin
	
	if a, ok := ctx.AppController.(NestPreparer); ok{
		a.NextPrepare()
	}
}

func (ctx *BaseController) MustLogin(){
	if !ctx.IsLogin{
		ctx.Abort500(syserror.NoUserError{})
	}
}

func (c *BaseController) GetMustString(key string, msg string) string{
	keyname := c.GetString(key, "")
	if len(keyname) == 0{
		c.Abort500(errors.New(msg))
	}

	return keyname
}

func (ctx *BaseController) Abort500(err error){
	ctx.Data["error"] = err
	ctx.Abort("500")
}

type H map[string]interface{}

type ResultJsonValue struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Action string `json:"action,omitempty"`
	Count int `json:"count,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOk(msg string, actions ...string){
	var action string
	if len(actions) > 0{
		action = actions[0]
	}

	ctx.Data["json"] = &ResultJsonValue{
		Code : 0,
		Msg : msg,
		Action : action,
		//Action : "/",
	}
	ctx.ServeJSON()
}

func (c *BaseController) JSONOkH(msg string, data H){
	data["code"] = 0
	data["msg"] = msg
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *BaseController) UUID() string{
	u, err := uuid.NewV4()
	if err != nil{
		c.Abort500(syserror.NewError("系统错误", err))
	}
	return u.String()
}