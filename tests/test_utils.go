package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/deyvidm/sms-backend/utils"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type expected struct {
	code int
	data map[string]interface{}
}

func toReader(obj interface{}) *bytes.Reader {
	objBytes, _ := json.Marshal(obj)
	return bytes.NewReader(objBytes)
}

func preTestSetup() {
	if router == nil {
		gin.SetMode(gin.ReleaseMode)
		router = utils.SetupRouter()
	}
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func getStepString(stepNumber int, stepName, errMessage string) string {
	return fmt.Sprintf("Error at step %d : %s : %s", stepNumber, stepName, errMessage)
}
