package controllers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"github.com/jinzhu/gorm"
	"golearn/lsyproject/liteblog/models"
	"golearn/lsyproject/liteblog/syserror"
)

type NoteController struct {
	BaseController
}

func (c *NoteController) NextPrepare(){
	//查询是否有登陆
	c.MustLogin()
	if c.User.Role != 0{
		c.Abort500(syserror.NewError("权限不足", nil))
	}
}

//note
// @router /new [get]
func (c *NoteController) NewPage(){
	c.Data["key"] = c.UUID()
	c.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (c *NoteController) EditPage(){
	key := c.Ctx.Input.Param(":key")

	note, err := models.QueryNoteByKeyAndUserId(key, int(c.User.ID))
	if err != nil{
		c.Abort500(syserror.NewError("文章不存在", err))
	}
	c.Data["note"] = note
	c.Data["key"] = key
	c.TplName = "note_new.html"
}

// @router /save/:key [post]
func (c *NoteController) Save(){
	key := c.Ctx.Input.Param(":key")
	editor := c.GetString("edit", "default")
	title := c.GetMustString("title", "标题不能为空！")
	content := c.GetMustString("content", "内容不能为空！")

	note, err := models.QueryNoteByKey(key)
	var n models.Note
	if err != nil{
		if err != gorm.ErrRecordNotFound{
			c.Abort500(syserror.NewError("保存失败", err))
		}
		
		n = models.Note{
			Key : key,
			Title : title,
			Content : content,
			UserID : int(c.User.ID),
		}
	}else{
		n = note
		n.Title = title
		n.Content = content					
	}
	n.Editor = editor

	//摘要
	n.Summary, _ = getSummary(content)

	if err = models.SaveNote(&n); err != nil{
		c.Abort500(syserror.NewError("保存失败", err))
	}
	c.JSONOk("保存成功", fmt.Sprintf("/details/%s", key))
}

// @router /del/:key [post]
func (c *NoteController) Del(){
	key := c.Ctx.Input.Param(":key")

    if err := models.DeleteNotByUserIdAndKey(key, int(c.User.ID)); err != nil{
		c.Abort500(syserror.NewError("删除失败", err))
	}

	c.JSONOk("删除成功", "/")
}


//获取摘要
func getSummary(html string) (string, error){
	var bufbytes bytes.Buffer

	//将字符串拼接
	if _, err := bufbytes.Write([]byte(html)); err != nil{
		return "", err
	}

	//变成Document格式
	doc, err := goquery.NewDocumentFromReader(&bufbytes)
	if err != nil{
		return "", nil
	}

	//取出body之后的数据
	htmlstr := doc.Find("body").Text()

	//当内容长度大于600，取前600为摘要
	if len(htmlstr) > 600{
		htmlstr = htmlstr[:600]
	}

	return htmlstr, nil
}