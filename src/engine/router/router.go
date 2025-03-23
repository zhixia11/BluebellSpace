package router

import (
	"BluebellSpace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *gin.Engine {
	route := gin.Default()
	route.Use(Cors())
	api := route.Group("/api")
	registerAPI(api)
	return route
}

func registerAPI(group *gin.RouterGroup) {
	registerSchedule(group.Group("/worker"))
	registerExplorer(group.Group("/explorer"))

}

// registerSchedule 调度组
func registerSchedule(group *gin.RouterGroup) {
	group.POST("/submit", service.SubmitWorker)
	group.GET("/progress", service.WorkerProgress)
	group.GET("/cancel", service.CancelWorker)
}

// registerExplorer 文件管理组
func registerExplorer(group *gin.RouterGroup) {
	group.POST("/delete", service.DeleteFile)
	group.GET("/info", service.FileInfo)
}

// Cors 中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
