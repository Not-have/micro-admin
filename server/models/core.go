package models

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	cfg, iniErr := ini.Load("./conf/app.ini")

	if iniErr != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
		os.Exit(1)
	}

	password := cfg.Section("mysql").Key("password").String()
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	app := cfg.Section("mysql").Key("app").String()

	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + app + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("连接DSN:", dsn) // 调试输出

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v\n", err) // 输出具体错误
		os.Exit(1)
	}
}
