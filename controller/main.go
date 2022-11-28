package controller

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Engine *gin.Engine
)

type ErrorMessage struct {
	Message string `json:"message"`
}

// CORS middleware
func middlewareCors() gin.HandlerFunc {
	cors_config := cors.DefaultConfig()
	cors_config.AllowAllOrigins = true
	cors_config.AllowWildcard = true
	// cors_config.AllowOrigins = []string{CORS_ORIGIN_URL}
	return cors.New(cors_config)
}

// Init initializes controller
func Init() {
	if Engine != nil {
		return
	}

	Engine = gin.Default()
	Engine.Use(middlewareCors())
	Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	Engine.POST("/api/v1/create", create_record)
	Engine.GET("/api/v1/get-content", get_content)
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func errorResp(c *gin.Context, error_code int, err error) {
	c.JSON(error_code, ErrorResponse{
		Message: err.Error(),
	})
}
