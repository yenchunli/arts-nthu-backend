package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/pkg/upload"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func (server *Server) UploadImage(ctx *gin.Context) {
	type request struct {
		image *multipart.FileHeader `form:image binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "miss data",
		})
		return
	}

	file, err := ctx.FormFile("image") // *Multipart.FileHeader
	if file.Size <= 0 {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	reader, err := file.Open() // io.Reader
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	buf, err := ioutil.ReadAll(reader) // bytes[]
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	client := upload.NewClient(os.Getenv("IMGUR_UPLOAD_TOKEN"), "https://api.imgur.com/3/upload")
	imgurUrl, _ := client.UploadImage(buf)

	ctx.JSON(200, gin.H{
		"url": imgurUrl,
	})
	return
}
