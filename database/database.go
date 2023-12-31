package database

import (
	"fmt"

	"github.com/SMRPcoder/markable/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance
var DB *gorm.DB

func Connetion() {
	dsn := `host=localhost user=postgres password=root dbname=markable port=5432 sslmode=disable TimeZone=Asia/Kolkata`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to connect database \n %v", err.Error()))
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.User{}, &models.Todo{}, &models.Taskmark{}, &models.Reminder{})

	Database = DbInstance{Db: db}
	DB = Database.Db

	// DB.Callback().Create().Before("gorm:before_create").Register("setCreatedAt",)

}
