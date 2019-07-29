package controllers

import (
	"errors"
	"strings"
	"golearn/lsyproject/liteblog/syserror"
	"golearn/lsyproject/liteblog/models"
	"golearn/lsyproject/liteblog/util"
)

type UserController struct {
	BaseController
}

//登陆
// @router /login [post]
func (c *UserController) Login(){
	//判断邮箱不为空
	email := c.GetMustString("email", "邮箱不能为空！")

	pwd := c.GetMustString("password", "密码不能为空！")

	var (
		user *models.User
		err error
	)

	//MD5加密
	md5pwd := util.Md5pwd(pwd)

	if user, err = models.QueryUserByEmailAndPassword(email, md5pwd); err != nil{
		c.Abort500(syserror.NewError("邮箱或密码不正确", err))
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"action": "/",
	}
	c.ServeJSON()
}

//注册
// @router /reg [post]
func (c *UserController) Reg(){
	name := c.GetMustString("name", "昵称不能为空")
	email := c.GetMustString("email", "邮箱不能为空")
	password := c.GetMustString("password", "密码不能为空")
	pwd2 := c.GetMustString("password2", "确认密码不能为空")

	if strings.Compare(password, pwd2) != 0{
		c.Abort500(errors.New("密码与确认密码必须一致"))
	}
	if u, err := models.QueryUserByName(name); err == nil && u != nil && u.ID != 0{
		c.Abort500(syserror.NewError("用户昵称已存在！", err))
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u != nil && u.ID != 0{
		c.Abort500(syserror.NewError("用户邮箱已存在！", err))
	}

	//MD5加密
	md5pwd := util.Md5pwd(password)

	if err := models.SaveUser(&models.User{
		Name : name,
		Email : email,
		Pwd : md5pwd,
		Avatar : "",
		Role : 1,
	}); err != nil{
		c.Abort500(syserror.NewError("用户注册失败", err))
	}
	c.JSONOk("注册成功", "/user")
}

//注销
// @router /logout [get]
func (c *UserController) Logout(){
	//检查是否已经登陆
	c.MustLogin()
	c.DelSession(SESSION_USER_KEY)
	c.Redirect("/", 302)
}