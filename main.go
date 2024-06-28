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
	// r.GET("/users/:id", userController.show)
	// r.POST("/users", userController.store)
	// r.PUT("/users/:id", userController.update)
	// r.DELETE("/users/:id", userController.destroy)

	
}