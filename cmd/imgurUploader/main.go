package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/pkg/upload"
	"io/ioutil"
	"os"
	"mime/multipart"
	"net/http"
)

func main() {

	r := gin.Default()
	r.POST("/api/v1/upload", func(c *gin.Context) {
		type request struct {
			image *multipart.FileHeader `form:image binding:"required"`
		}
		var req request
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "miss data",
			})
			return
		}


		file, err := c.FormFile("image") 	// *Multipart.FileHeader
		if file.Size <=0 {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		reader, err := file.Open()			// io.Reader
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		buf, err := ioutil.ReadAll(reader)	// bytes[]
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		client := upload.NewClient(os.Getenv("IMGUR_UPLOAD_TOKEN"), "https://api.imgur.com/3/upload")
		imgurUrl, _ := client.UploadImage(buf)


		c.JSON(200, gin.H{
			"url": imgurUrl,
		})
		return
	})
	r.Run() 
	
}