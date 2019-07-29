package controllers

import (
	"github.com/astaxie/beego/logs"
	"golearn/lsyproject/liteblog/syserror"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if (c.IsAjax()) {
		c.jsonerror(syserror.Error404{})
	} else {
		c.Data["content"] = "非法访问"
	}
}

func (c *ErrorController) Error500(){
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok{
		err = syserror.NewError("未知错误", nil)
	}

	serr, ok := err.(syserror.Error)
	if !ok{
		serr = syserror.NewError(err.Error(), nil)
	}

	if serr.ReasonError() != nil{
		logs.Info(serr.Error(), serr.ReasonError())
	}

	if c.IsAjax(){
		c.jsonerror(serr)
	}else{
		c.Data["content"] = serr.Error()
	}
}

func (c *ErrorController) jsonerror(serr syserror.Error){
	c.Ctx.Output.Status = 200
	c.Data["json"] = &ResultJsonValue{
		Code: serr.Code(),
		Msg:  serr.Error(),
	}
	c.ServeJSON()
}