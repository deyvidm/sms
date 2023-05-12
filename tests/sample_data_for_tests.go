package tests

import "github.com/deyvidm/sms-backend/types"

var Users = []types.LoginData{{
	Username: "testUser",
	Password: "hunter2",
}, {
	Username: "testUser2",
	Password: "hunter3",
}}

var Contacts = []types.NewContactData{
	{
		FirstName: "Florian",
		LastName:  "Degas",
		Phone:     "+11234567890",
	},
	{
		FirstName: "Sneed",
		LastName:  "Feedenseed",
		Phone:     "+11234567891",
	},
	{
		FirstName: "Michael",
		LastName:  "Pichaelson",
		Phone:     "+11234567892",
	},
}
