package utils

import (
	"log"
	"model"
	"os"

	"github.com/imshuai/serverchan"
	//mysql数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	lg "github.com/imshuai/lightlog"
)

var (
	//DatabaseCfg 运行时储存数据库相关配置
	DatabaseCfg *DatabaseConfig
	sc          *serverchan.ServerChan
)

func init() {
	lg.SetConsole(false)
	lg.SetLevel(lg.DEBUG)
	lg.SetPrefix("GoTyping")
	lg.SetRollingDaily("logs", "GoTyping.log")
	var err error
	DatabaseCfg, err = parseDatabaseConfig()
	if err != nil {
		LogFatal("connot parse database config")
		os.Exit(1)
	}
	databaseInit()
	serverchanInit()
	loggerInit()
}

//DB 数据库引擎
var DB *xorm.Engine

//databaseInit 初始化数据库连接
func databaseInit() error {
	var err error
	DB, err = xorm.NewEngine("mysql", DatabaseCfg.String())
	if err != nil {
		log.Fatalln("connect to database fail with error:", err)
	}
	DB.SetMapper(core.GonicMapper{})
	DB.SetMaxIdleConns(DatabaseCfg.DBIdleConnectionNum)
	DB.SetMaxOpenConns(DatabaseCfg.DBMaxConnectionNum)
	return DB.Sync2(&model.Artical{}, &model.Comment{}, &model.Category{})
}
