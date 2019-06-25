package models

import "cooky-go/models"

type RoleMenu struct {
	RoleId int
	MenuId int
}

func (RoleMenu) TableName() string {
	return "t_role_menu"
}

func AddRoleMenu(roleId int, menuIds []int) {
	DeleteRoleMenuByRoleId(roleId)
	for _, menuId := range menuIds {
		models.DB.Create(&RoleMenu{roleId, menuId})
	}
}

func DeleteRoleMenuByRoleId(roleId int) {
	models.DB.Delete(RoleMenu{}, "role_id=?", roleId)
}

func DeleteRoleMenuByMenuId(roleId int) {

}
