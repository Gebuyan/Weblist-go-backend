package routers

import (
	"Weblist-go-backend/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//
	v1Group := r.Group("v1")
	{
		//增加待办事项
		v1Group.POST("/todo", controller.CreateATodo)
		//删除待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
		//修改待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//查看待办事项
		//查看所以待办事项
		v1Group.GET("/todo", controller.GetTodolist)
		//查看某个待办事项
		v1Group.GET("/todo/:id", controller.GetATodo)
	}
	return r
}
