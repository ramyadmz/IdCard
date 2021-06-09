package main

import (
	"fmt"

	_ "idcheck/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var router *gin.Engine

// @title National ID API
// @version 1.0
// @description This is a National ID API Documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://ramyadmz.github.io/CV/
// @contact.email ramyadmz@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	fmt.Println("...")

	router = gin.Default()

	url := ginSwagger.URL("127.0.0.1:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initialize the routes
	initRoute()

	// Start serving the application
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
