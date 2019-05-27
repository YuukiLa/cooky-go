package models

import "cooky-go/models"

type Dept struct {
	models.Model
	DeptId int `gorm:"primary_key" json:"userId"`
	DeptName string `json:"deptName"`
	ParentId int `json:"parentId"`
	Remark string `json:"remark"`
}

func (Dept) TableName() string {
	return "t_dept"
}