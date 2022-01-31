package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	port := "13306"
	dbname := "graphdb"
	username := "root"
	password := "localhost"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate()
}
