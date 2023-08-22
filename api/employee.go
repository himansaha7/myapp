package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var employee_map map[string]employee

func saveEmployee(c echo.Context) error {
	//u := new(employee)
	u := employee{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	if employee_map == nil {
		employee_map = make(map[string]employee)
	}
	employee_map[u.Id] = u
	fmt.Println(employee_map)

	return c.JSON(http.StatusCreated, u)
}

func getEmployee(c echo.Context) error {
	id := c.Param("id")
	if emp, ok := employee_map[id]; ok {
		return c.JSON(http.StatusOK, emp)
	}
	return c.JSON(http.StatusNotFound, nil)
}

func deleteEmployee(c echo.Context) error {
	id := c.Param("id")
	if _, ok := employee_map[id]; ok {
		delete(employee_map, id)
		return c.JSON(http.StatusOK, "deleted")
	}
	return c.JSON(http.StatusNotFound, nil)
}
