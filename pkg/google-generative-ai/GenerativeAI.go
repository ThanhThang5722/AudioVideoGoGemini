package models

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var client *genai.Client
var model *genai.GenerativeModel

func GetClient() *genai.Client {
	return client
}

func CreateNewClient(ctx *context.Context, apiKey string) {
	option := option.WithAPIKey(apiKey)
	cl, err := genai.NewClient(*ctx, option)
	client = cl
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
	}
}

func ConnectGemini(client *genai.Client) {
	model = client.GenerativeModel("gemini-1.5-pro")
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
}

func GetModelInstance() *genai.GenerativeModel {
	return model
}

func UploadToGemini(path, mimeType string, client *genai.Client, ctx context.Context) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	options := genai.UploadFileOptions{
		DisplayName: path,
		MIMEType:    mimeType,
	}
	fileData, err := client.UploadFile(ctx, "", file, &options)
	if err != nil {
		log.Fatalf("Error uploading file: %v\n", err)
	}

	log.Printf("Uploaded file %s as: %s\n", fileData.DisplayName, fileData.URI)
	return fileData.URI
}
