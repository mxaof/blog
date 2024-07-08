package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mxaof_blog/dao/global"
)

var MysqlConn *gorm.DB

func MysqlInit() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.GlobalConfig.Mysql.Username, global.GlobalConfig.Mysql.Password, global.GlobalConfig.Mysql.Host, global.GlobalConfig.Mysql.Port, global.GlobalConfig.Mysql.DbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
		return err
	}
	MysqlConn = db
	return nil
}
