package controllers

import (
	"golearn/lsyproject/liteblog/syserror"
	"golearn/lsyproject/liteblog/models"

)

//Index API
type IndexController struct{
	BaseController
}

// @router / [get]
func (this *IndexController) Index() {
	limit := 4
	//page
	page, err := this.GetInt("page", 1)

	if err != nil || page <= 0{
		page = 1
	}

	title := this.GetString("title")

	//得到页面数据
	notes, err := models.QueryNoteByPage(page, limit, title)
	if err != nil{
		this.Abort500(err)
	}
	this.Data["notes"] = notes

	//总页数
	count, err := models.QueryNoteCount(title)

	if err != nil{
		this.Abort500(err)
	}

	totalpage := count / limit

	if count % limit != 0{
		totalpage += 1
	}

	this.Data["totalpage"] = totalpage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /user [get]
func(this *IndexController) GetUser(){
	this.TplName = "user.html"
}

// @router /reg [get]
func(this *IndexController) GetReg(){
	this.TplName = "reg.html"
}

// @router /details/:key [get]
func(this *IndexController) GetDetails(){
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)

	if err != nil{
		this.Abort500(syserror.NewError("文章不存在", err))
	}

	this.Data["note"] = note
	this.TplName = "details.html"
}

// @router /comment/:key [get]
func(this *IndexController) GetComment(){
	key := this.Ctx.Input.Param(":key")

	note, err := models.QueryNoteByKey(key)

	if err != nil{
		this.Abort500(syserror.NewError("文章不存在", err))
	}

	this.Data["note"] = note
	this.TplName = "comment.html"
}