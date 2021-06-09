package main

import "github.com/gin-gonic/gin"

// CORSMiddleware godoc
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initRoute() {
	// Group persons related routes together
	router.Use(CORSMiddleware())

	personsRoutes := router.Group("/ID")
	{
		// Handle GET requests at /ID/validation/:id
		personsRoutes.GET("/validation/:id", CheckID)

	}
}
