package models

import (
	"cooky-go/models"
	"github.com/jinzhu/gorm"
	"time"
)

type Menu struct {
	models.Model
	MenuId        int    `gorm:"primary_key" json:"menuId"`
	Permission    string `json:"permission"`
	MenuName      string `json:"menuName"`
	ComponentName string `json:"componentName"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	MenuType      int    `json:"menuType"`
	ParentId      int    `json:"parentId"`
}

func (Menu) TableName() string {
	return "t_menu"
}

func SelectAllMenu() (menus []Menu) {
	models.DB.Find(&menus)
	return
}

func (menu *Menu) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ct", time.Now())
	_ = scope.SetColumn("mt", time.Now())
	return nil
}

func (menu *Menu) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("mt", time.Now())
	return nil
}
