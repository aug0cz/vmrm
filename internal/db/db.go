package db

import (
	"github.com/aug0cz/vmrm/internal/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	dbFileName = "test.db"
)

func init() {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&user.User{}, &user.Department{})
	if err != nil {
		panic(err)
	}
	DB = db
}
