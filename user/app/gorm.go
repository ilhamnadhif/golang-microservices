package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/config"
	"user/helper"
)

func InitGorm(config config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.HostPort, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.FatalIfError(err)

	return db
}
