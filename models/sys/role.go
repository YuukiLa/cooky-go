package models

import (
	"cooky-go/models"
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	models.Model
	RoleId   int    `gorm:"primary_key" json:"roleId"`
	RoleName string `json:"roleName"`
	Remark   string `json:"remark"`
}

func (Role) TableName() string {
	return "t_role"
}

func SelectAllRole() (roles []Role) {
	models.DB.Find(&roles)
	return
}

func AddRole(role Role) bool {
	models.DB.Create(&role)
	return true
}

func (role *Role) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ct", time.Now())
	_ = scope.SetColumn("mt", time.Now())
	return nil
}

func (role *Role) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("mt", time.Now())
	return nil
}
