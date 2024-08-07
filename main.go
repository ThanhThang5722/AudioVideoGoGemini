package main

//AIzaSyDUuELGwaUNJ-MS_eKWkKB9fNH9llgBsPM
import (
	GenAI "GolangGemini/pkg/google-generative-ai"
	"GolangGemini/routes"
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	// Connect Gemini key
	apiKey := os.Getenv("GeminiAPI")

	GenAI.CreateNewClient(&ctx, apiKey)
	client := GenAI.GetClient()
	defer client.Close()
	GenAI.ConnectGemini(client)

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // Max file size upload
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ping successful",
			})
		})
	}
	//ROUTER DEFINE
	routes.VideoRouter(api)
	routes.MessageRouter(api)
	router.Run()
}
