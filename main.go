package main

import (
	"Weblist-go-backend/dao"
	"Weblist-go-backend/models"
	"Weblist-go-backend/routers"
)

func main() {
	//创建数据库
	//sql: create database weblist;
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//模型绑定(绑定到表上，如果表不存在则创建)
	dao.DB.AutoMigrate(&models.Todo{}) //表名：todos

	//初始化路由
	r := routers.SetupRouter()
	//启动
	r.Run(":9000")
}
