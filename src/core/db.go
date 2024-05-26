package core

import (
	"flyblog/src/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	// 初始化数据库连接
	dbConfig := config.DbConfig
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// 循环dbconfig
	for _, dbConfig := range dbConfig.Databases {
		if dbConfig.Drive == "mysql" {
			// 初始化数据库连接
			dsn = fmt.Sprintf(dsn, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
			break
		}
	}
	fmt.Println(dsn)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	Db = db
}
