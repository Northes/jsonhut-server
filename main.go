package main

import (
	"fmt"
	"jsonhutapi/dao"
	"jsonhutapi/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		//fmt.Println(err.Error())
		panic(err.Error())
	}

	// 注册路由
	r := routers.SetupRouter()

	err = r.Run(":8081")
	if err != nil {
		fmt.Println(err.Error())
	}
}
