package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//定义db变量，变量名开头小写，只能在同包下访问
	db *gorm.DB
)

func init(){

	var err error

	db, err = gorm.Open("mysql", "root:lin111@(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil{
		panic("failed to connect database.")
	}
	
	// 自动同步表结构
	db.AutoMigrate(&User{},&Note{})

	var count int
    // Model(&User{})查询用户表, Count(&count) 将用户表的数据赋值给count字段。
    if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
        //新增
        db.Create(&User{Name: "admin",
            //邮箱
            Email: "admin@qq.com",
            //密码
            Pwd: "123123",
            //头像地址
            Avatar: "/static/images/info-img.png",
            //角色 管理员
            Role: 0,
        })
	}
	
	//defer db.Close()
}
