package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/deyvidm/sms-backend/types"
	utils "github.com/deyvidm/sms-backend/utils"
)

func TestUserRegisterLoginFlow(t *testing.T) {
	cleanupDB := utils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	user1 := types.LoginData{
		Username: "testUser",
		Password: "hunter2",
	}
	user2 := types.LoginData{
		Username: "testUser",
		Password: "hunter3", // different password from user1, to simulate a bad login
	}

	steps := []struct {
		name string
		path string
		body types.LoginData
		exp  types.ExpectedResponse
	}{
		{name: "missing user", path: "/user/login", body: user1,
			exp: types.ExpectedResponse{Code: 400, ReturnBody: map[string]interface{}{
				"status": types.StatusFailed,
				"data":   "incorrect login details",
			}}},
		{name: "register new user", path: "/user/register", body: user1,
			exp: types.ExpectedResponse{Code: 200, ReturnBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   fmt.Sprintf("welcome %s!", user1.Username),
			}}},
		{name: "log in wrong user", path: "/user/login", body: user2,
			exp: types.ExpectedResponse{Code: 400, ReturnBody: map[string]interface{}{
				"status": types.StatusFailed,
				"data":   "incorrect login details",
			}}},
		{name: "log in correct user", path: "/user/login", body: user1,
			exp: types.ExpectedResponse{Code: 200, ReturnBody: map[string]interface{}{
				"status": types.StatusSuccess,
				// "data":   " a successful login returns a unique token that we can't reproduce (except through mocking 🤮), so we leave data nil
			}}},
	}

	for i, s := range steps {
		w := performRequest(router, http.MethodPost, s.path, toReader(s.body))
		s.exp.Compare(r, w, getStepString(i, s.name))
	}
}
