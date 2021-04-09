package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/pkg/upload"
)

func main() {

	r := gin.Default()
	r.POST("/api/v1/upload", func(c *gin.Context) {

		image, _ := c.FormFile("image")

		file, _ := image.Open()

		client := upload.NewClient(os.GetEnv("IMGUR_UPLOAD_TOKEN"), "https://api.imgur.com/3/upload")
		imgurUrl, _ := client.UploadImage(file)


		c.JSON(200, gin.H{
			"url": imgurUrl,
		})
	})
	r.Run() 
	
}