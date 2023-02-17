package sql

import (
	"douSheng/Const"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := Const.User + ":" + Const.Pass + "@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接数据库
	if err != nil {
		panic("failed to connect database" + fmt.Sprintf("%s", err))
	}
}

func getDB() *gorm.DB {
	return db
}

func getNewDB() *gorm.DB {
	dsn := Const.User + ":" + Const.Pass + "@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接数据库
	if err != nil {
		panic("failed to connect database" + fmt.Sprintf("%s", err))
	}
	return db
}
