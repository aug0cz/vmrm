package datacenter

import "gorm.io/gorm"

type DC struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Addr string
}

type DcColumn struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex"`
	Size           uint
	DatacenterName string
	Datacenter     DC `gorm:"references:Name;foreignKey:DatacenterName"`
}

type Cabinet struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex"`
	Index      uint
	Size       uint
	ColumnName string
	Column     DcColumn `gorm:"references:Name;foreignKey:ColumnName"`
}

type CabinetPosition struct {
	gorm.Model
	Name        string
	Index       uint
	CabinetName string
	Cabinet     Cabinet `gorm:"references:Name;foreignKey:CabinetName"`
}
