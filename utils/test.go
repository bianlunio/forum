package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReturnStruct struct {
	Page      int                      `json:"page"`
	Size      int                      `json:"size"`
	Total     int                      `json:"total"`
	List      []map[string]interface{} `json:"list"`
	Detail    map[string]interface{}   `json:"detail"`
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

//func GetTestDB(session *mgo.Session) *mgo.Database {
//	return session.DB("forum_test")
//}
//
