package model

import (
	"database/sql"
	"log"

	// 数据库驱动注册及初始化
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB     //全局变量

func Init() error {

	var err error

	//这行代码的作用就是初始化一个sql.DB对象
	DB ,err = sql.Open("mysql", viper.GetString("mysql.url"))
	if nil != err {
		return err
	}

	//设置最大超时时间
	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	//建立链接
	err = DB.Ping()
	if nil != err{
		return err
	}else{
		log.Println("Mysql Startup Normal!")
	}
	return nil
}
