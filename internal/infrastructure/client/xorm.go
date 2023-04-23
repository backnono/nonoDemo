// Package client @Author nono.he 2023/4/20 16:51:00
package client

import (
	"fmt"
	"log"

	//"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"nonoDemo/pkg/config"
	"xorm.io/xorm"
)

const (
	DBTypeMysql   = "mysql"
	DBTypeSqlite3 = "sqlite3"
)

var DBConnectFormat = map[string]string{
	DBTypeMysql:   "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&interpolateParams=true&parseTime=true&loc=Local",
	DBTypeSqlite3: "./training.db",
}

func NewXorm(config config.Config) *xorm.Engine {
	format, ok := DBConnectFormat[DBTypeMysql]
	if !ok {
		panic("暂不支持的数据库类型")
	}
	connStr := fmt.Sprintf(
		format, config.Database.User, config.Database.Password,
		config.Database.Host, config.Database.Port, config.Database.Name)
	x, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		log.Panicf("连接数据库失败：%s", err)
	}
	if config.Database.ShowSql {
		x.ShowSQL(true)
	}
	return x
}
