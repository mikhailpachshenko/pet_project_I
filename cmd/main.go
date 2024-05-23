package main

import (
	nH "github.com/mikhailpachshenko/pet_project_I.git/handlers"
	m "github.com/mikhailpachshenko/pet_project_I.git/models"

	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := m.NewMemoryStorage()
	handler := nH.NewHandler(memoryStorage)

	router := gin.Default() // создаём объект типа *gin.Engine

	router.POST("/employee", handler.CreateEmployee)       // создание сотрудника
	router.GET("/employee/:id", handler.GetEmployee)       // получение информации о сотруднике
	router.PUT("/employee/:id", handler.UpdateEmployee)    // редактирование информации о сотруднике
	router.DELETE("/employee/:id", handler.DeleteEmployee) // удаление сотрудника

	router.Run() // запуск HTTP сервера
}
