package handlers

import (
	models "GolangGemini/pkg/google-generative-ai"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

type MessageFromUSer struct {
	MSG string `json:"msg"`
}

func SendToGemini(ctx *gin.Context) {
	gemini := models.GetModelInstance()
	chat := gemini.StartChat()

	var mess MessageFromUSer
	if err := ctx.ShouldBind(&mess); err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := chat.SendMessage(ctx, genai.Text(mess.MSG))
	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusExpectationFailed, gin.H{
		"response": fmt.Sprintln(r.Candidates[0].Content.Parts[0]),
	})
}

func SendVideoAndTextToGemini(ctx *gin.Context) {
	client := models.GetClient()
	gemini := models.GetModelInstance()
	form, _ := ctx.MultipartForm()
	video := form.File["video"]
	//prompt := form.Value["prompt"]

	SaveVideoToDisk(video, ctx)

	filePath, err := GetFilePath(video[0].Filename)
	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	opts := genai.UploadFileOptions{DisplayName: "Video"}
	response, err := client.UploadFile(ctx, "", f, &opts)
	if err != nil {
		log.Fatal(err)
	}

	// View the response.
	var file *genai.File = response
	fmt.Printf("Uploaded file %s as: %q\n", file.DisplayName, file.URI)
	response, err = client.GetFile(ctx, file.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Poll GetFile() on a set interval (10 seconds here) to
	// check file state.
	for response.State == genai.FileStateProcessing {
		fmt.Print(".")
		// Sleep for 10 seconds
		time.Sleep(10 * time.Second)

		// Fetch the file from the API again.
		response, err = client.GetFile(ctx, file.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	prompt := []genai.Part{
		genai.FileData{URI: response.URI},
		genai.Text("Hãy tạo phụ đề cho nội dung của Video trên kèm thêm timeline"),
	}

	// Generate content using the prompt.
	resp, err := gemini.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"result": resp.Candidates[0].Content.Parts[0],
	})
	/*
		    // Handle the response of generated text.
		    for _, c := range resp.Candidates {
		        if c.Content != nil {
					res = append(res, *c.Content)
		            fmt.Println(*c.Content)
		        }
		    }*/
}

func SoundToText(ctx *gin.Context) {
	client := models.GetClient()
	gemini := models.GetModelInstance()
	form, _ := ctx.MultipartForm()
	audio := form.File["audio"]
	//prompt := form.Value["prompt"]

	SaveVideoToDisk(audio, ctx)

	filePath, err := GetFilePath(audio[0].Filename)
	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	opts := genai.UploadFileOptions{DisplayName: "Audio"}
	response, err := client.UploadFile(ctx, "", f, &opts)
	if err != nil {
		log.Fatal(err)
	}

	// View the response.
	var file *genai.File = response
	fmt.Printf("Uploaded file %s as: %q\n", file.DisplayName, file.URI)
	response, err = client.GetFile(ctx, file.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Poll GetFile() on a set interval (10 seconds here) to
	// check file state.
	for response.State == genai.FileStateProcessing {
		fmt.Print(".")
		// Sleep for 10 seconds
		time.Sleep(10 * time.Second)

		// Fetch the file from the API again.
		response, err = client.GetFile(ctx, file.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	prompt := []genai.Part{
		genai.FileData{URI: response.URI},
		genai.Text("Hãy tạo văn bản theo nội dung từ audio trên"),
	}

	// Generate content using the prompt.
	resp, err := gemini.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"result": resp.Candidates[0].Content.Parts[0],
	})
	/*
		    // Handle the response of generated text.
		    for _, c := range resp.Candidates {
		        if c.Content != nil {
					res = append(res, *c.Content)
		            fmt.Println(*c.Content)
		        }
		    }*/
}
