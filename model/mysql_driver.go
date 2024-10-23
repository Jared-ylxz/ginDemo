package model

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(conf *ini.File) *gorm.DB {
	mysqlName := conf.Section("mysql").Key("db_name").String()
	mysqlUser := conf.Section("mysql").Key("db_user").String()
	mysqlPwd := conf.Section("mysql").Key("db_pwd").Value()
	mysqlHost := conf.Section("mysql").Key("db_host").String()
	mysqlPort := conf.Section("mysql").Key("db_port").String()
	mysqlCharset := conf.Section("mysql").Key("db_charset").String()
	mysqlPwd = mysqlPwd + "#"
	fmt.Println(mysqlPwd)
	dsn := mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlName + "?charset=" + mysqlCharset
	fmt.Println(dsn)
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return DB
}