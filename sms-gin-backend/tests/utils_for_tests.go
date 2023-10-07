package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/deyvidm/sms-backend/routes"
	"github.com/deyvidm/sms-backend/types"
	"github.com/deyvidm/sms-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// This type is used for our tests, and expresses what sort of response we expect from http handlers
type ExpectedResponse struct {
	Code         int
	ResponseBody map[string]interface{}
}

// testName is useful for "stepped" tests, where you'd like to know at which step
func (er *ExpectedResponse) Compare(r *require.Assertions, w *httptest.ResponseRecorder, stepName string) {
	er.CompareHTTPCode(r, w.Code, stepName)
	er.CompareResponseBody(r, w.Body, stepName)
}

func (er *ExpectedResponse) CompareHTTPCode(r *require.Assertions, respCode int, stepName string) {
	r.Equal(er.Code, respCode, fmt.Sprintf("error: mismatched HTTP Codes | %s", stepName))
}

func (er *ExpectedResponse) CompareResponseBody(r *require.Assertions, respBody *bytes.Buffer, stepName string) {
	var responseBody map[string]interface{}
	_ = json.Unmarshal(respBody.Bytes(), &responseBody)

	_, exists := er.ResponseBody["data"]
	if !exists { // sometimes we don't want to (or can't) match the exact value of data
		er.ResponseBody["data"] = nil
		responseBody["data"] = nil
	}
	r.Equal(er.ResponseBody, responseBody, fmt.Sprintf("error: mismatched response data | %s", stepName))
}

type TestStep struct {
	name   string
	method string
	path   string
	body   interface{}
	exp    ExpectedResponse
}

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

func authUser(user types.NewUser) string {
	performRequest(router, http.MethodPost, routes.UserRegister, toReader(user))
	w := performRequest(router, http.MethodPost, routes.UserLogin, toReader(user))
	return getTokenFromLoginResponse(w)
}
