package server

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/yenchunli/arts-nthu-backend/store"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRouterMethod(t *testing.T) {
	store := store.NewMockStore()
	router := NewRouter(store)

	w := performRequest(router, http.MethodGet, "/api/v1/exhibitions")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, http.MethodGet, "/api/v1/exhibitions/1")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, http.MethodPost, "/api/v1/exhibitions")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, http.MethodPut, "/api/v1/exhibitions/1")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, http.MethodDelete, "/api/v1/exhibitions/1")
	assert.Equal(t, http.StatusOK, w.Code)
}

/*
func TestExhibitionService(t *testing.T) {
	store := NewMockStore()
	svc := NewExhibitionSvc(store)
}
*/
