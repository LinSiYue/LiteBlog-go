package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Key string `gorm:"unique:not null"`
	UserID int
	User User
	Title string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Source  string `gorm:"type:text" json:"source"`
	Editor  string `gorm:"varchar(40)'" `
	Files   string `gorm:"type:text"`
	Visit int `gorm:"type:varchar(800)"`
	Praise int `gorm:"default:0"`
}

func QueryNoteByKey(key string) (notes Note, err error){
	return notes, db.Where("`key` = ?", key).Take(&notes).Error
}

func QueryNoteByKeyAndUserId(key string, userId int) (notes Note, err error){
	return notes, db.Where("`key` = ? and `user_id` = ?", key, userId).Take(&notes).Error
}

func SaveNote(note *Note) error{
	return db.Save(note).Error
}

func QueryNoteByPage(page, limit int, title string) (note []*Note, err error){
	//"%title%"
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).Order("updated_at desc").Offset((page - 1) * limit).Limit(limit).Find(&note).Error
}

func QueryNoteCount(title string) (count int, err error){
	return count, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Count(&count).Error
}

func DeleteNotByUserIdAndKey(key string, userid int) error{
	return db.Delete(&Note{},"`key` = ? and `user_id` = ?", key, userid).Error
}