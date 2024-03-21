package main

import (
	cntrl "latFramework/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Router *gin.Engine

func main() {
	Router = gin.Default()

	Router.Use(cntrl.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		// v1.GET("/users/:id", controllers.GetUserDetail)
		v1.GET("/users/", cntrl.GetUser)
		// v1.POST("/login/", controllers.Login)
		// v1.PUT("/users/:id", controllers.UpdateUser)
		// v1.POST("/users", controllers.PostUser)
	}

	Router.Run(":8888")
}
