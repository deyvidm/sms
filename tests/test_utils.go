package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	setuputils "github.com/deyvidm/sms-backend/setupUtils"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type expected struct {
	code    int
	message string
	data    interface{}
}

func toReader(obj interface{}) *bytes.Reader {
	objBytes, _ := json.Marshal(obj)
	return bytes.NewReader(objBytes)
}

func preTestSetup() {
	if router == nil {
		gin.SetMode(gin.ReleaseMode)
		router = setuputils.SetupRouter()
	}
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
