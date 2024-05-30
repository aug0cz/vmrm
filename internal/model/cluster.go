package model

import "gorm.io/gorm"

type Cluster struct {
	gorm.Model
	Name string
}
