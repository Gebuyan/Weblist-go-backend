package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/weblist?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//sql: create database weblist;
	//连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型绑定(绑定到表上，如果表不存在则创建)
	DB.AutoMigrate(&Todo{})

	r := gin.Default()

	//
	v1Group := r.Group("v1")
	{
		//增加待办事项
		v1Group.POST("/todo", func(context *gin.Context) {
			var todo Todo
			if err := context.ShouldBindJSON(&todo); err == nil {
				fmt.Println("todo is :", todo)
				DB.Create(&todo)
				context.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}

		})
		//删除待办事项
		v1Group.DELETE("/todo/:id", func(context *gin.Context) {

			id := context.Param("id")
			if err := DB.Delete(&Todo{}, id); err == nil {
				context.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
			}
		})
		//修改待办事项
		v1Group.PUT("/todo/:id", func(context *gin.Context) {
			id := context.Param("id")
			var todo Todo
			err := context.ShouldBindJSON(&todo)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}
			if err := DB.Model(&Todo{}).Where("id = ?", id).Update("status", todo.Status); err == nil {
				context.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
			}

		})
		//查看待办事项
		//查看所以待办事项
		v1Group.GET("/todo", func(context *gin.Context) {
			var todos []Todo
			if err := DB.Find(&todos); err.Error == nil {
				context.JSON(http.StatusOK, todos)
			} else {
				context.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
			}
		})
		//查看某个待办事项
		v1Group.GET("/todo/:id", func(context *gin.Context) {
			var todo Todo
			id := context.Param("id")
			if err := DB.Where("id = ?", id).Find(&todo); err == nil {
				context.JSON(http.StatusOK, todo)
			} else {
				context.JSON(http.StatusBadRequest, gin.H{
					"err": err,
				})
			}
		})
	}

	r.Run(":9000")
}
