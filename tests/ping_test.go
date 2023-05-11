package tests

import (
	"testing"

	"github.com/deyvidm/sms-backend/types"
	utils "github.com/deyvidm/sms-backend/utils"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	cleanupDB := utils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	expected := types.ExpectedResponse{
		Code: 200,
		ResponseBody: map[string]interface{}{
			"status": types.StatusSuccess,
			"data":   "pong",
		},
	}
	w := performRequest(router, "GET", "/ping", nil)
	expected.Compare(r, w, "ping/pong test :)")
}
