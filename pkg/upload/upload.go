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

func (client *Client) UploadImage(image io.Reader) (string, error){
	var buf = new(bytes.Buffer)
    writer := multipart.NewWriter(buf)

    part, _ := writer.CreateFormFile("image", "dont care about name")
    io.Copy(part, image)

    writer.Close()
    req, _ := http.NewRequest("POST", clinet.uploadApiUrl, buf)
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