package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReturnStruct struct {
	Pagination map[string]interface{}   `json:"pagination"`
	List       []map[string]interface{} `json:"list"`
}

func ListResponse2Dict(b []byte) (v ReturnStruct) {
	err := json.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}
	return v
}

func Response2Dict(b []byte) (v map[string]interface{}) {
	err := json.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}
	return v
}

func TestRequest(router *gin.Engine, method string, url string, body string) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)
	w := httptest.NewRecorder()
	var req *http.Request
	req, _ = http.NewRequest(method, url, reader)
	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}
	router.ServeHTTP(w, req)
	return w
}
