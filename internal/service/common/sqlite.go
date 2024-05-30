package common

import (
	"context"
	"log"

	"github.com/aug0cz/vmrm/internal/model"
	"github.com/aug0cz/vmrm/internal/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbFileName = "test.db"
)

type ContextKey = string

func NewSqliteConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database.")
	}

	err = db.AutoMigrate(&model.Cluster{})
	if err != nil {
		log.Fatalln("db auto migrate error", err)
	}
	return db
}

func NewSqliteConnectWithContext() context.Context {
	db := NewSqliteConnect()
	bgCtx := context.Background()
	ctx := context.WithValue(bgCtx, util.DBContextKey, db)
	return ctx
}
