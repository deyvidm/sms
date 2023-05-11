package tests

import (
	"encoding/json"
	"testing"

	setuputils "github.com/deyvidm/sms-backend/setupUtils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	cleanupDB := setuputils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	w := performRequest(router, "GET", "/ping", nil)
	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	value, exists := response["data"]
	r.Nil(err)
	r.True(exists)
	r.Equal("pong", value)
}
