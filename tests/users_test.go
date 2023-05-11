package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/deyvidm/sms-backend/controllers"
	"github.com/deyvidm/sms-backend/types"
	"github.com/stretchr/testify/require"

	utils "github.com/deyvidm/sms-backend/utils"
)

func TestUserRegisterLoginFlow(t *testing.T) {
	cleanupDB := utils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	user := controllers.LoginData{
		Username: "testUser",
		Password: "hunter2",
	}

	steps := []struct {
		name string
		path string
		body controllers.LoginData
		exp  expected
	}{
		{name: "missing user", path: "/user/login", body: user,
			exp: expected{400, map[string]interface{}{
				"status": types.StatusFailed,
				"data":   "incorrect login details",
			}}},
		{name: "register new user", path: "/user/register", body: user,
			exp: expected{203, map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   fmt.Sprintf("welcome %s!", user.Username),
			}}},
		// {name: "log in wrong user"},
		// {name: "log in correct user"},
	}

	for i, s := range steps {
		w := performRequest(router, http.MethodPost, s.path, toReader(s.body))

		r.Equal(s.exp.code, w.Code, getStepString(i, s.name, "mismatched HTTP Codes"))
		var response map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &response)
		r.Equal(s.exp.data, response, getStepString(i, s.name, "mismatched response data"))
	}
}
