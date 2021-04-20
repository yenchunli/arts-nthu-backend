package upload

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

// Test if upload a nil image, it should return error
func TestUploadNilImage(t *testing.T) {

	client := NewClient("", "")

	_, err := client.UploadImage(nil)

	if err == nil {
		t.Error("UploadImage() should have an error")
		t.Fail()
	}
}

// Test if upload real image, it should return success(200)
func TestUploadRealImage(t *testing.T) {
	token := os.Getenv("IMGUR_UPLOAD_TOKEN")
	if token == "" {
		t.Skip("IMGUR_UPLOAD_TOKEN is not set.")
	}

	client := NewClient(os.Getenv("IMGUR_UPLOAD_TOKEN"), "https://api.imgur.com/3/upload")

	// Read File to byte
	file, err := ioutil.ReadFile("logo.png")
	if err != nil {
		t.Skip("Can't read logo.png for test")
	}

	url, err := client.UploadImage(file)
	if err != nil {
		t.Errorf("UploadImage() failed with error: %v", err)
		t.Fail()
	}
	if matched, _ := regexp.MatchString(`https://i.imgur.com/`, url); !matched {
		t.Error("UploadImage() did not return imgur url")
		t.Fail()
	}

}
