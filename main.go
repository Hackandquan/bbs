package main

import (
	"bbs/controller"
	"bbs/pkg/config"
	"fmt"

	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

func init() {
	// 全局配置
	conf := config.Init("./config.yaml")
	fmt.Printf("%+v", conf)
	// gorm 配置
	gormConf := &gorm.Config{}
	// 连接数据库
	if err := simple.OpenDB(conf.MySqlUrl, gormConf, 10, 20); err != nil {
		panic("gorm:" + err.Error())
	}
}

func main() {
	controller.Router()
}
