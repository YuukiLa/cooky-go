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

//func (menu *Menu) UnmarshalJSON(b []byte) error {
//	menu.MenuId,_ = com.StrTo(string(b)).Int()
//	return nil
//}

func SelectAllMenu() (menus []Menu) {
	models.DB.Find(&menus)
	return
}

func SelectMenuByIds(ids []int) (menus []Menu) {
	models.DB.Where("menu_id IN (?) and url != ''", ids).Find(&menus)
	return
}

func AddMenu(menu Menu) bool {
	models.DB.Save(&menu)
	return true
}

func EditMenu(menu Menu) bool {
	if menu.ParentId == 0 && menu.MenuName == "" {
		models.DB.Model(&Menu{}).Where("menu_id=?", menu.MenuId).Update("parent_id", menu.ParentId)
		return true
	}
	models.DB.Model(&Menu{}).Where("menu_id=?", menu.MenuId).Update(menu)
	return true
}

func FindMenusByUserId(userId int) (menus []Menu) {
	models.DB.Debug().Raw("SELECT m.* FROM t_menu m LEFT JOIN t_role_menu rm ON rm.`menu_id`=m.`menu_id` LEFT JOIN t_user_role ur ON ur.`role_id` = rm.`role_id` WHERE ur.`user_id` =?", userId).Scan(&menus)
	return
}

func DeleteMenu(menuId int) {
	models.DB.Delete(Menu{}, "menu_id=?", menuId)
	models.DB.Model(&Menu{}).Where("parent_id=?", menuId).Update("parent_id", 0)
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
