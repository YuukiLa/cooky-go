package models

import (
	"cooky-go/models"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	models.Model
	UserId    int    `gorm:"primary_key" json:"userId"`
	Username  string `json:"username" binding:"required" `
	Password  string `json:"password" binding:"required"`
	DeptId    int    `json:"deptId"`
	DeptName  string `json:"deptName" gorm:"-"'`
	Sex       int    `json:"sex"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Described string `json:"described"`
	Status    int    `json:"status"`
}

/**
指定表名
*/
func (User) TableName() string {
	return "t_user"
}

func SelectUser(pageNum int, pageSize int, maps interface{}) (users []User) {
	models.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)
	return
}

func FindByUsername(username string) (user User) {
	models.DB.Where("username=?", username).First(&user)
	return
}

func AddUser(user User) bool {
	models.DB.Create(&user)
	return true
}

func GetUserTotal(maps interface{}) (count int) {
	models.DB.Model(&User{}).Where(maps).Count(&count)

	return
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ct", time.Now())
	_ = scope.SetColumn("mt", time.Now())
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("mt", time.Now())
	return nil
}
