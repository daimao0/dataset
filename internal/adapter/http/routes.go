package http

import (
	"dataset/internal/adapter/http/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// RegisterRoutes register all routes
func RegisterRoutes(engine *gin.Engine) {
	// new controller
	datasetController := controller.NewDatasetController()
	// handle cors config
	config := cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源
		AllowMethods:     []string{"*"}, // 允许的 HTTP 方法
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Datasource-Id"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 是否允许发送凭据
		MaxAge:           12 * time.Hour, // 预检请求的有效期
	}
	engine.Use(cors.New(config))
	// GET /api/v1/database/list
	group := engine.Group("/api/")
	v1 := group.Group("/v1")
	//datasource route group
	datasetGroup := v1.Group("/dataset")
	{
		datasetGroup.GET("/get/:id", datasetController.GetDataset)
		datasetGroup.POST("/create/json", datasetController.CreateDatasetByJSON)
		datasetGroup.POST("/create/file", datasetController.CreateDatasetByFile)
	}
}
