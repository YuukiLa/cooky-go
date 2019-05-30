package models

import "cooky-go/models"

type Menu struct {
	models.Model
	MenuId        int    `gorm:"primary_key" json:"menuId"`
	MenuName      string `json:"menuName"`
	CompenentName string `json:"compenentName"`
	Url           string `json:"url"`
	MenuType      int    `json:"menuType"`
	ParentId      int    `json:"parentId"`
}

func (Menu) TableName() string {
	return "t_menu"
}
