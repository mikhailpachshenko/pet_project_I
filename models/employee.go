package model

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
	Salary int    `json:"salary"`
}
