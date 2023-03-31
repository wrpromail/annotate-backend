package pkg

import (
	"github.com/gin-gonic/gin"
	"mymlops/annotate-helper/pkg/service"
)

func RegisterFileController(router *gin.RouterGroup) {
	fileGroup := router.Group("/file")
	fileGroup.GET("/", service.ListFile)
	fileGroup.GET("/:id", service.GetFile)
	fileGroup.GET("/:id/ontology", service.GetFileOntology)

	// 整个文件浏览、标注与报告
	fileGroup.GET("/:id/browse")
	fileGroup.POST("/:id/browse/annotate")
	fileGroup.POST("/:id/browse/report")

	fileGroup.GET("/:id/line", service.GetFileLineCount)
	fileGroup.GET("/:id/line/:number", service.GetFileLine)
	fileGroup.POST("/:id/line/:number/annotate", service.FileLineAnnotate)
	fileGroup.POST("/:id/line/:number/report", service.FileLineReport)
}
