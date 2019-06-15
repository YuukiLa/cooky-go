package models

import (
	"cooky-go/models"
	"github.com/jinzhu/gorm"
	"time"
)

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

func EditDept(dept Dept) bool {
	models.DB.Model(&Dept{}).Where("dept_id = ?", dept.DeptId).Update(dept)
	return true
}

func (dept *Dept) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ct", time.Now())
	_ = scope.SetColumn("mt", time.Now())
	return nil
}

func (dept *Dept) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("mt", time.Now())
	return nil
}
