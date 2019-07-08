package models

import "cooky-go/models"

type UserRole struct {
	UserId int
	RoleId int
}

/**
指定表名
*/
func (UserRole) TableName() string {
	return "t_user_role"
}

func AddUserRole(userId int, roleIds []int) {
	DeleteUserRoleByUserId(userId)
	for _, roleId := range roleIds {
		models.DB.Create(&UserRole{userId, roleId})
	}
}

func DeleteUserRoleByUserId(userId int) {
	models.DB.Delete(UserRole{}, "user_id=?", userId)
}
