package handlers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Video *multipart.FileHeader `form:"video"`
}

func GetFilePath(FileName string) (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	dest := dir + `\\assets\\` + FileName
	return dest, nil
}

func SaveVideoToDisk(files []*multipart.FileHeader, ctx *gin.Context) {
	for _, file := range files {
		log.Println(file.Filename)
		dest, err := GetFilePath(file.Filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		// Upload the file to specific dst.
		ctx.SaveUploadedFile(file, dest)
	}
}

func PostVideo(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]
	SaveVideoToDisk(files, ctx)
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success!",
	})
}
