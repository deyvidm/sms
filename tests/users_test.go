package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/deyvidm/sms-backend/controllers"
	setuputils "github.com/deyvidm/sms-backend/setupUtils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRegisterLoginFlow(t *testing.T) {
	cleanupDB := setuputils.SetupDB("")
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
			exp: expected{400, "", map[string]interface{}{
				"error": "incorrect login details",
			}}},
		{name: "register new user", path: "/users/register", body: user,
			exp: expected{200, "welcome user!", nil}},
		// {name: "log in wrong user"},
		// {name: "log in correct user"},
	}

	for i, s := range steps {
		w := performRequest(router, http.MethodPost, s.path, toReader(s.body))
		r.Equal(w.Code, s.exp.code)

		var response map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &response)

		value, exists := response["message"]
		if s.exp.message != "" {
			if !exists {
				t.Fatalf("test %d : %s : missing message. expected |%s| got nothin", i, s.name, s.exp.message)
			}
			if value != s.exp.message {
				t.Fatalf("test %d : %s : wrong message. expected |%s| got |%s|", i, s.name, s.exp.message, value)
			}
			assert.True(t, exists)
		}

		r.Equal(s.exp.data, response)
	}
}
