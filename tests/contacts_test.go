package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/deyvidm/sms-backend/routes"
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

	steps := []TestStep{
		{name: "add contact", method: http.MethodPost, path: routes.NewContact, body: contact,
			exp: ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   contactResponse,
			}}},
		{name: "fetch all contacts", method: http.MethodGet, path: routes.AllContacts, body: contact,
			exp: ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   []interface{}{contactResponse},
			}}},
	}
	token := authUser(user)
	for i, s := range steps {
		w := performAuthRequest(router, s.method, s.path, token, toReader(s.body))
		s.exp.Compare(r, w, getStepString(i, s.name))
	}
}
