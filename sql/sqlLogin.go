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
	dsn := Const.User + ":" + Const.Pass + "@tcp(" + Const.Host + ":" + Const.Addr + ")/" + Const.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接数据库
	if err != nil {
		panic("failed to connect database" + fmt.Sprintf("%s", err))
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
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
