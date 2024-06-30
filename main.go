package main

import (
	"awesomeAPI/controllers/labelcontroller"
	"awesomeAPI/controllers/usercontroller"
	"awesomeAPI/middleware"
	"awesomeAPI/models"

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
	
	users := r.Group("/users")
	{
		users.GET("", usercontroller.Index)
		users.POST("", usercontroller.Create)
		users.GET("/:id", usercontroller.Show)
		users.PATCH("/:id", usercontroller.Update)
		users.DELETE("/:id", usercontroller.Delete)
	}
	auth := r.Group("/auth")
	{
		auth.POST("/login", usercontroller.Login)
		auth.GET("/profile", middleware.RequireAuth, usercontroller.Profile)
	}

	labels := r.Group("/labels")
	{
		labels.GET("/:id", labelcontroller.ShowLabel)
	}

	r.Run()
}