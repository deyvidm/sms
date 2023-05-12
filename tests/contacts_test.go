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

func TestAddContactNoAuth(t *testing.T) {
	cleanupDB := utils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	exp := ExpectedResponse{Code: 401, ResponseBody: map[string]interface{}{
		"status": types.StatusFailed,
		"data":   "you need to authenticate for this request",
	}}
	w := performAuthRequest(router, http.MethodPost, routes.NewContact, "", toReader(Contacts[0]))
	exp.Compare(r, w, "add contact without auth")
}

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
		{name: "fetch all contacts", method: http.MethodGet, path: routes.AllContacts,
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

func TestAddMultipleContactsFlow(t *testing.T) {
	cleanupDB := utils.SetupDB("")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	token := authUser(Users[0])
	for i := range Contacts {
		exp := ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
			"status": types.StatusSuccess,
			"data":   utils.ObjToJSONObj(Contacts[i]),
		}}

		w := performAuthRequest(router, http.MethodPost, routes.NewContact, token, toReader(Contacts[i]))
		exp.Compare(r, w, getStepString(i, "adding contact"))
	}

	exp := ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
		"status": types.StatusSuccess,
		"data":   utils.ObjToJSONObj(Contacts),
	}}

	w := performAuthRequest(router, http.MethodGet, routes.AllContacts, token, nil)
	exp.Compare(r, w, "comparing response of all contacts")

}
