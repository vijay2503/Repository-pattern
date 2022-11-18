package router

import (
	_ "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func Routing() *echo.Echo {
	e := echo.New()
	e.POST("/set/employee-detail", CreateEmployee)
	e.GET("/All-employee", GetAllEmployee)
	e.PUT("Update", UpdateEmail)
	e.DELETE("Delete/employee", DeleteDetails)
	return e
}
