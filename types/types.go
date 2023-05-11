package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

const StatusFailed = "failed"
const StatusSuccess = "success"

type LoginData struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=255"` // 3 is a holy number
	Password string `json:"password" binding:"required,alphanum,min=6,max=255"` // min 6 for brcypt hash
}

type NewContactData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=3,max=50"`
	LastName  string `json:"last_name" binding:"required,alpha,min=3,max=50"`
	Phone     string `json:"phone" binding:"required,e164"` // e164 is the standard +11234567890
}

type NewEventData struct {
	Ttile string `json:"title" binding:"required,max=255"`
}

// This type is used for our tests, and expresses what sort of response we expect from http handlers
type ExpectedResponse struct {
	Code       int
	ReturnBody map[string]interface{}
}

func (er *ExpectedResponse) Compare(r *require.Assertions, w *httptest.ResponseRecorder, errString string) {
	er.CompareHTTPCode(r, w.Code, errString)
	er.CompareResponseBody(r, w.Body, errString)
}

func (er *ExpectedResponse) CompareHTTPCode(r *require.Assertions, respCode int, errString string) {
	r.Equal(er.Code, respCode, fmt.Sprintf("error: mismatched HTTP Codes | %s", errString))
}

func (er *ExpectedResponse) CompareResponseBody(r *require.Assertions, respBody *bytes.Buffer, errString string) {
	var responseBody map[string]interface{}
	_ = json.Unmarshal(respBody.Bytes(), &responseBody)

	_, exists := er.ReturnBody["data"]
	if !exists { // sometimes we don't want to (or can't) match the exact value of data
		er.ReturnBody["data"] = nil
		responseBody["data"] = nil
	}
	r.Equal(er.ReturnBody, responseBody, fmt.Sprintf("error: mismatched response data | %s", errString))
}
