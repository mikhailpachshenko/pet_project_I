package main

import (
	nH "github.com/mikhailpachshenko/pet_project_I.git/handler"
	m "github.com/mikhailpachshenko/pet_project_I.git/models"

	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := m.NewMemoryStorage()
	handler := nH.NewHandler(memoryStorage)

	router := gin.Default() // создаём объект типа *gin.Engine

	router.POST("/employee")       // создание сотрудника
	router.GET("/employee/:id")    // получение информации о сотруднике
	router.PUT("/employee/:id")    // редактирование информации о сотруднике
	router.DELETE("/employee/:id") // удаление сотрудника

	router.Run() // запуск HTTP сервера
}
