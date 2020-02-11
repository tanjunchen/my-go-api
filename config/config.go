package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// LogInfo 初始化日志配置
func LogInfo()   {
	file := "./" + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(file,os.O_RDWR| os.O_CREATE| os.O_APPEND, 0766)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
}

// Config viper解析配置文件
func Config() error  {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
// Init 读取初始化配置文件
func Init() error {

	// 初始化配置文件
	if err := Config(); err != nil {
		return err
	}

	// 初始化日志包
	LogInfo()
	return nil
}
