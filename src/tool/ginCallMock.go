package tool

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"net/http/httptest"
)

func GinCall(para any, handlers ...gin.HandlerFunc) *httptest.ResponseRecorder {
	// def para
	marshal, _ := json.Marshal(para)
	reader := bytes.NewReader(marshal)
	req, _ := http.NewRequest("POST", "/", reader)

	// def server
	router := gin.Default()
	router.POST("/", handlers...)
	// call
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}
