package router

import (
	"encoding/json"
	"fmt"
	"go-postgres/driver"
	"go-postgres/model"
	"go-postgres/repository"
	"go-postgres/repository/repoimpl"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var userRepo repository.UserRepo

//function Init
func init() {
	if err := godotenv.Load("dbConfig.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	db := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	userRepo = repoimpl.NewUserRepo(db.SQL)
}

func GetAllEmployee(c echo.Context) error {
	employee := userRepo.Select()
	AllEmployee, err := json.Marshal(employee)
	if err != nil {
		return c.String(http.StatusBadRequest, "RECORD NOT FOUND")
	} else {
		jsonResponse := fmt.Sprintf("successfully Get Details%s", string(AllEmployee))
		return c.String(http.StatusOK, jsonResponse)
	}
}
func CreateEmployee(c echo.Context) error {
	create := model.User{}
	err := c.Bind(&create)
	if err != nil {
		return err
	}
	Err := userRepo.Create(create)
	if Err != nil {
		return Err
	}
	jsonValue, err := json.Marshal(create)

	if err != nil {
		return c.String(http.StatusBadRequest, "Not SET the employee detail")
	} else {
		return c.String(http.StatusOK, string(jsonValue))
	}

}
func UpdateEmail(c echo.Context) error {
	id := c.QueryParam("userId")
	userId, _ := strconv.Atoi(id)
	update := model.User{}
	err := c.Bind(&update)
	if err != nil {
		return err
	}
	err = userRepo.Update(&userId, &update)
	if err != nil {
		return c.String(http.StatusBadRequest, "Email-Updation is Fail")
	}
	return c.String(http.StatusOK, "Successfully updated")
}
func DeleteDetails(c echo.Context) error {
	id := c.QueryParam("userId")
	userId, _ := strconv.Atoi(id)
	err := userRepo.Delete(userId)
	if err != nil {
		return c.String(http.StatusBadRequest, "Deleted operation is Un-Success")
	}
	return c.String(http.StatusOK, "Success-fully Dleted the Employee Detail")
}
