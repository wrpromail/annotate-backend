package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mymlops/annotate-helper/pkg"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	engine := gin.New()
	engine.Use(CORS())
	v1Group := engine.Group("/api/v1")

	pkg.RegisterFileController(v1Group)

	if err := engine.Run(":6789"); err != nil {
		log.Fatal(err)
	}
}
