package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/deyvidm/sms-backend/routes"
	"github.com/deyvidm/sms-backend/types"
	utils "github.com/deyvidm/sms-backend/utils"
)

// this test is a series of steps
// each step must pass in order for the next one to succeed
// i will use this pattern throughout the testing files
//
// each step is an HTTP request to the /users/ endpoints
// this test simulates a user that:
// 1. attempts to log in without an account, and fails
// 2. registers
// 3. logs in with bad credentials
// 4. logs in with correct credentials
func TestUserRegisterLoginFlow(t *testing.T) {
	cleanupDB := utils.SetupDB("../.env", "")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	steps := []TestStep{
		{name: "try to log in with non-existent user", method: http.MethodPost, path: routes.UserLogin, body: LoginUser[0],
			exp: ExpectedResponse{Code: 400, ResponseBody: map[string]interface{}{
				"status": types.StatusFailed,
				"data":   "incorrect login details",
			}}},
		{name: "register new user", method: http.MethodPost, path: routes.UserRegister, body: NewUser[0],
			exp: ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   fmt.Sprintf("welcome %s!", LoginUser[0].Username),
			}}},
		{name: "log in user", method: http.MethodPost, path: routes.UserLogin, body: LoginUser[0],
			exp: ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				// "data":  a successful login returns a unique token that we can't reproduce (except through mocking 🤮), so we leave data nil
				// the Compare() method will only compare non-nil fields in the Resopnse Body
			}}},
	}

	for i, s := range steps {
		w := performRequest(router, s.method, s.path, toReader(s.body))
		s.exp.Compare(r, w, getStepString(i, s.name))
	}
}
