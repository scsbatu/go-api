package models

import (
	"fmt"
	"github.com/scsbatu/go-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var Layout = "2006-01-02 15:04:05"
var DOBLayout = "2006-01-02"
var Location *time.Location

func Init() error {
	Location = time.Local
	return initializeDBConn()
}

func initializeDBConn() (err error) {
	connectionStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.DataBase.Username,
		config.Config.DataBase.Password,
		config.Config.DataBase.Host,
		config.Config.DataBase.Port,
		config.Config.DataBase.Name,
	)
	db, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	return err
}
