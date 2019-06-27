package models

import (
	"cooky-go/models"
	"github.com/Unknwon/com"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type User struct {
	models.Model
	UserId   int    `gorm:"primary_key" json:"userId"`
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required"`
	DeptId   int    `json:"deptId"`
	DeptName string `json:"deptName" gorm:"-"'`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Remark   string `json:"remark"`
	Status   int    `json:"status"`
	RoleIds  []int  `json:"roleIds" gorm:"-"`
	RoleStr  string `json:"-" gorm:"-"`
}

/**
指定表名
*/
func (User) TableName() string {
	return "t_user"
}

func SelectUser(pageNum int, pageSize int, maps interface{}) (users []User) {
	models.DB.Raw("SELECT u.*,d.`dept_name`,GROUP_CONCAT(ur.`role_id`) role_str FROM t_user u LEFT JOIN t_dept d ON d.`dept_id`=u.`dept_id` LEFT JOIN t_user_role ur ON ur.`user_id`=u.`user_id`").Where(maps).Offset(pageNum).Limit(pageSize).Scan(&users)
	for i := 0; i < len(users); i++ {
		ids := strings.Split(users[i].RoleStr, ",")
		users[i].RoleIds = make([]int, len(ids))
		for j := 0; j < len(ids); j++ {
			users[i].RoleIds[j] = com.StrTo(ids[j]).MustInt()
		}
	}
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
