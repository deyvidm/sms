package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/deyvidm/sms-backend/types"
	utils "github.com/deyvidm/sms-backend/utils"
	"github.com/stretchr/testify/require"
)

func TestNewContactFlow(t *testing.T) {
	cleanupDB := utils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	user := types.LoginData{
		Username: "testUser",
		Password: "hunter2",
	}

	contact := types.NewContactData{
		FirstName: "Florian",
		LastName:  "Degas",
		Phone:     "+11234567890",
	}
	var contactResponse map[string]interface{}
	b, _ := json.Marshal(contact)
	json.Unmarshal(b, &contactResponse)

	steps := []struct {
		name   string
		method string
		path   string
		body   interface{}
		exp    types.ExpectedResponse
	}{
		{name: "register new user", method: http.MethodPost, path: "/user/register", body: user,
			exp: types.ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   fmt.Sprintf("welcome %s!", user.Username),
			}}},
		{name: "log in correct user", method: http.MethodPost, path: "/user/login", body: user,
			exp: types.ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				// see users_test.go for why data is nil
			}}},
		{name: "add contact", method: http.MethodPost, path: "/api/contacts/new", body: contact,
			exp: types.ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   contactResponse,
			}}},
	}

	i := 0 // register user
	performRequest(router, steps[i].method, steps[i].path, toReader(steps[i].body))
	i++ // login & capture token
	w := performRequest(router, steps[i].method, steps[i].path, toReader(steps[i].body))
	token := getTokenFromLoginResponse(w)
	i++
	w = performAuthRequest(router, steps[i].method, steps[i].path, token, toReader(steps[i].body))
	steps[i].exp.Compare(r, w, getStepString(i, steps[i].name))

}
