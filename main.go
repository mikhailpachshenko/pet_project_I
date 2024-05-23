package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // создаём объект типа *gin.Engine

	router.POST("/employee")       // создание сотрудника
	router.GET("/employee/:id")    // получение информации о сотруднике
	router.PUT("/employee/:id")    // редактирование информации о сотруднике
	router.DELETE("/employee/:id") // удаление сотрудника

	router.Run() // запуск HTTP сервера
}
