package models

import (
	"time"
)

type User struct{
	Model
	Name string `gorm:"unique_index" json:"name"`
	Email string `gorm:"unique_index" json:"email"`
	Avatar string `json:"avatar"`
	Pwd string `json:"-"`
	Role int `gorm:"default:0" json:"role"` //1 管理员 0 普通用户
}

//查询帐号密码
func QueryUserByEmailAndPassword(email, password string) (*User, error){
	//这就是gorm的魅力，一句代码进行语句查询
	var users User
	//查询数据库随机一条数据
	if err := db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&users).Error; err != nil{
		return nil, err
	}
	return &users, nil
}

func QueryUserByName(name string) (*User, error){
	var users User

	if err := db.Where("name = ?", name).Take(&users).Error; err != nil{
		return nil, err
	}
	return &users, nil
}

func QueryUserByEmail(email string) (*User, error){
	var users User

	if err := db.Where("email = ?", email).Take(&users).Error; err != nil{
		return nil, err
	}
	return &users, nil
}

func SaveUser(user *User) error{
	return db.Create(user).Error
}

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}