package main

import (
	"fmt"
	"jsonhutapi/config"
	"jsonhutapi/dao"
	"jsonhutapi/models"
	"jsonhutapi/routers"
)

func main() {
	config.LoadConfig()
	// 连接MySQL数据库
	err := dao.InitMySQL()
	if err != nil {
		//fmt.Println(err.Error())
		panic(err.Error())
	}
	// 连接Redis
	dao.InitRedis()
	// 自动迁移
	err = dao.DB.AutoMigrate(&models.Json{})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 注册路由
	r := routers.SetupRouter()

	err = r.Run(config.App.Port)
	if err != nil {
		fmt.Println(err.Error())
	}
}
