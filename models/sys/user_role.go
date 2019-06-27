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

func FindRoleIdByUserId(userId int) (roles []int) {
	models.DB.Select("role_id").Where("user_id = ?", userId).Find(&roles)
	return
}
