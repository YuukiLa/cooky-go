package models

import (
	"cooky-go/models"
	"github.com/Unknwon/com"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Role struct {
	models.Model
	RoleId   int    `gorm:"primary_key" json:"roleId"`
	RoleName string `json:"roleName"`
	Remark   string `json:"remark"`
	MenuIds  []int  `json:"menuIds" gorm:"-"`
	MenuStr  string `json:"-" gorm:"-"`
}

func (Role) TableName() string {
	return "t_role"
}

func SelectAllRole() (roles []Role) {
	models.DB.Raw("SELECT r.*,GROUP_CONCAT(rm.`menu_id`) menu_str FROM t_role r LEFT JOIN t_role_menu rm ON r.`role_id`=rm.`role_id` GROUP BY r.`role_id`").Scan(&roles)

	for i := 0; i < len(roles); i++ {
		ids := strings.Split(roles[i].MenuStr, ",")
		roles[i].MenuIds = make([]int, len(ids))
		for j := 0; j < len(ids); j++ {
			roles[i].MenuIds[j] = com.StrTo(ids[j]).MustInt()
		}
	}
	return
}

func AddRole(role Role) bool {
	menus := SelectMenuByIds(role.MenuIds)
	//models.DB.Create(&role)
	AddPolicy(role, menus)
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
