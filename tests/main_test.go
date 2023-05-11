package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	setuputils "github.com/deyvidm/sms-backend/setupUtils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func preTestSetup() {
	fmt.Println("\t Router nil? ", router == nil)
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

func TestPing(t *testing.T) {
	cleanupDB := setuputils.SetupDB("")
	defer cleanupDB()
	preTestSetup()

	w := performRequest(router, "GET", "/ping", nil)
	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	value, exists := response["data"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "pong", value)
}
