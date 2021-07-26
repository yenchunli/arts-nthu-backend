package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
