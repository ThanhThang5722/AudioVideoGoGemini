package routes

import (
	handler "GolangGemini/handlers"

	"github.com/gin-gonic/gin"
)

func VideoRouter(r *gin.RouterGroup) {
	videoRouter := r.Group("/video")
	{
		videoRouter.POST("/upload", handler.PostVideo)
		videoRouter.POST("/gemini", handler.SendVideoAndTextToGemini)
		videoRouter.POST("/audioGemini", handler.SoundToText)
	}
}
