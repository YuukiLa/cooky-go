package models

import "cooky-go/models"

type Dept struct {
	models.Model
	DeptId   int    `gorm:"primary_key" json:"deptId"`
	DeptName string `json:"deptName"`
	ParentId int    `json:"parentId"`
	Remark   string `json:"remark"`
}

func (Dept) TableName() string {
	return "t_dept"
}

func SelectDept(pageNum int, pageSize int, maps interface{}) (depts []Dept) {
	models.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&depts)
	return
}

func SelectAllDept() (depts []Dept) {
	models.DB.Find(&depts)
	return
}

func AddDept(dept Dept) bool {
	models.DB.Create(&dept)
	return true
}
