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

	// (文本)文件某一行的浏览、标注与报告
	fileGroup.GET("/:id/line", service.GetFileLineCount)
	fileGroup.GET("/:id/line/:number", service.GetFileLine)
	fileGroup.POST("/:id/line/:number/annotate", service.FileLineAnnotate)
	fileGroup.POST("/:id/line/:number/report", service.FileLineReport)
}

// RegisterJobController 从任务的维度
//func RegisterJobController(router *gin.RouterGroup) {
//	jobGroup := router.Group("/job")
//	// 查看自己的任务
//	jobGroup.GET("/", service.ListJob)
//	jobGroup.POST("/", service.CreateJob)
//	jobGroup.PUT("/", service.UpdateJob)
//	jobGroup.GET("/:id", service.GetJob)
//	jobGroup.DELETE("/:id", service.DeleteJob)
//}
