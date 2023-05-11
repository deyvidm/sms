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

func newRequest(method, path, token string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}

func doRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performAuthRequest(r http.Handler, method, path, token string, body io.Reader) *httptest.ResponseRecorder {
	return doRequest(r, newRequest(method, path, token, body))
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	return doRequest(r, newRequest(method, path, "", body))
}

func getTokenFromLoginResponse(resp *httptest.ResponseRecorder) string {
	var responseBody map[string]string
	_ = json.Unmarshal(resp.Body.Bytes(), &responseBody)
	return responseBody["data"]
}

func getStepString(stepNumber int, stepName string) string {
	return fmt.Sprintf("Error at step %d : \"%s\" ", stepNumber, stepName)
}
