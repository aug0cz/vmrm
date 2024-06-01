package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string
	Email          string
	NameEn         string
	DepartmentName string
	Department     Department `gorm:"references:ShortName;foreignKey:DepartmentName"`
}

type Department struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	ShortName   string `gorm:"uniqueIndex"`
	Description string
}
