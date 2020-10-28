package handlers

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Router returns the api resources
func SetupRouter(service Service) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	mazeAPI := router.Group("/api/v1/merchants")
	mazeAPI.POST("/", service.CreateHandler)
	mazeAPI.DELETE("/", service.DeleteIDHandler)
	mazeAPI.PATCH("/", service.UpdateHandler)
	mazeAPI.GET("/", service.GetHandler)
	mazeAPI.GET("/", service.GetIDHandler)
	return router
}
