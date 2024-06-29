package main

import (
	"awesomeAPI/models"
	"awesomeAPI/controllers/usercontroller"
	// "database/sql"
	// "errors"
	// "fmt"
	// "log"
	// "net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	r := gin.Default()
	models.ConnectDatabase()
	
	r.GET("/users", usercontroller.Index)
	r.POST("/users", usercontroller.Create)

	r.Run()
}