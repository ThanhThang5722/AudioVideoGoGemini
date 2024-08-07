package routes

import (
	handler "GolangGemini/handlers"

	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.RouterGroup) {
	videoRouter := r.Group("/message")
	{
		videoRouter.POST("/", handler.SendToGemini)
	}
}
