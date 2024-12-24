package controller

import (
	"Weblist-go-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url -> controller -> logic -> models
*/

func CreateATodo(context *gin.Context) {
	//1.获取参数
	var todo models.Todo
	context.BindJSON(&todo)
	//2.存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(context *gin.Context) {

	id := context.Param("id")
	if err := models.DeleteATodo(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"id": "deleted",
		})
	}
}

func UpdateATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{
			"err": "无效的id",
		})
	}
	var todo models.Todo

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
	if err := models.UpdateATodoById(id, &todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, todo)
	}

}

func GetTodolist(context *gin.Context) {
	todolist, err := models.GetTodolist()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, todolist)
	}
}

func GetATodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := models.GetATodo(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})

	} else {
		context.JSON(http.StatusOK, todo)
	}
}
