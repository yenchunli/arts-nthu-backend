package upload

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"io"
	"mime/multipart"
	"encoding/json"
	"errors"
)

type Client struct {
	token string
	uploadApiUrl string
}

func NewClient(token string, uploadApiUrl string) *Client{
	return &Client{token: token, uploadApiUrl: uploadApiUrl}
}

func (client *Client) UploadImage(image []byte) (string, error){

	if image == nil {
		return "", errors.New("No Image")
	}
	var buf = new(bytes.Buffer)
    writer := multipart.NewWriter(buf)

    part, _ := writer.CreateFormFile("image", "filename")

	imgReader := bytes.NewReader(image)
    io.Copy(part, imgReader)

    writer.Close()
    req, _ := http.NewRequest("POST", client.uploadApiUrl, buf)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Authorization", "Bearer "+client.token)

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

	dec := json.NewDecoder(bytes.NewReader(body))
	var img imageInfoDataWrapper
	if err := dec.Decode(&img); err != nil {
		return "", errors.New("Fail to decode")
	}

	if !img.Success {
		return "", errors.New("Fail")
	}
    
	return img.Data.Link, nil
}