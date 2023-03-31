package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mymlops/annotate-helper/pkg/service"
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

	fileGroup := v1Group.Group("/file")
	fileGroup.GET("/", service.ListFile)
	fileGroup.GET("/:id", service.GetFile)
	fileGroup.GET("/:id/ontology", service.GetFileOntology)
	fileGroup.GET("/:id/line", service.GetFileLineCount)
	fileGroup.GET("/:id/line/:number", service.GetFileLine)
	fileGroup.POST("/:id/line/:number/annotate", service.FileLineAnnotate)
	if err := engine.Run(":6789"); err != nil {
		log.Fatal(err)
	}
}
