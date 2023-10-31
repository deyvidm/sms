package tests

import (
	"net/http"
	"testing"

	"github.com/deyvidm/sms/web-server/models"
	"github.com/deyvidm/sms/web-server/routes"
	"github.com/deyvidm/sms/web-server/types"
	utils "github.com/deyvidm/sms/web-server/utils"
	"github.com/stretchr/testify/require"
)

func TestAddContactNoAuth(t *testing.T) {
	cleanupDB := utils.SetupDB("../.env", "")
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
	cleanupDB := utils.SetupDB("../.env", "")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	contactResponse := models.APIContact{
		FirstName: Contacts[0].FirstName,
		LastName:  Contacts[0].LastName,
		Phone:     Contacts[0].Phone,
		ID:        2,
	}

	steps := []TestStep{
		{name: "add contact", method: http.MethodPost, path: routes.NewContact, body: Contacts[0],
			exp: ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   utils.ObjToJSONObj(contactResponse),
			}}},
		{name: "fetch all contacts", method: http.MethodGet, path: routes.AllContacts,
			exp: ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
				"status": types.StatusSuccess,
				"data":   []interface{}{utils.ObjToJSONObj(contactResponse)},
			}}},
	}
	token := authUser(NewUser[0])
	for i, s := range steps {
		w := performAuthRequest(router, s.method, s.path, token, toReader(s.body))
		s.exp.Compare(r, w, getStepString(i, s.name))
	}
}

func TestAddMultipleContactsFlow(t *testing.T) {
	cleanupDB := utils.SetupDB("../.env", "")
	defer cleanupDB()
	preTestSetup()
	r := require.New(t)

	token := authUser(NewUser[0])
	var contactResponses []models.APIContact
	for i := range Contacts {
		cr := models.APIContact{
			FirstName: Contacts[i].FirstName,
			LastName:  Contacts[i].LastName,
			Phone:     Contacts[i].Phone,
			ID:        uint(i + 2),
		}
		contactResponses = append(contactResponses, cr)
		exp := ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
			"status": types.StatusSuccess,
			"data":   utils.ObjToJSONObj(cr),
		}}

		w := performAuthRequest(router, http.MethodPost, routes.NewContact, token, toReader(Contacts[i]))
		exp.Compare(r, w, getStepString(i, "adding contact"))
	}

	exp := ExpectedResponse{Code: 200, ResponseBody: map[string]interface{}{
		"status": types.StatusSuccess,
		"data":   utils.ObjToJSONObj(contactResponses),
	}}

	w := performAuthRequest(router, http.MethodGet, routes.AllContacts, token, nil)
	exp.Compare(r, w, "comparing response of all contacts")

}
