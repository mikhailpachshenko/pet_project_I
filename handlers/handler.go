package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	m "github.com/mikhailpachshenko/pet_project_I.git/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage m.Storage
}

func NewHandler(storage m.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee m.Employee
	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("FAILED: to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, m.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	h.storage.Insert(&employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("FAILED: to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, m.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("FAILED: to get employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, m.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     employee.ID,
		"Name":   employee.Name,
		"Age":    employee.Age,
		"Sex":    employee.Sex,
		"Salary": employee.Salary,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("FAILED: to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, m.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	var employee m.Employee
	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("FAILED: to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, m.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Update(id, employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee,
	})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("FAILED: to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, m.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Delete(id)
	c.JSON(http.StatusOK, "SUCCESS: employee has been deleted")
}
