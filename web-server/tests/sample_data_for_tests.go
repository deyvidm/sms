package tests

import "github.com/deyvidm/sms/web-server/types"

var FirstLast = []types.FirstLastName{
	{
		FirstName: "Florian",
		LastName:  "Degas",
	}, {
		FirstName: "Sneed",
		LastName:  "Feedenseed",
	}, {
		FirstName: "Michael",
		LastName:  "Pichaelson",
	},
}

var LoginUser = []types.LoginUser{
	{
		Username: "testUser",
		Password: "hunter2",
	}, {
		Username: "testUser2",
		Password: "hunter3",
	}, {
		Username: "testUser3",
		Password: "hunter3",
	},
}

var NewUser = []types.NewUser{
	{
		FirstLastName: FirstLast[0],
		LoginUser:     LoginUser[0],
	},
	{
		FirstLastName: FirstLast[1],
		LoginUser:     LoginUser[1],
	},
	{
		FirstLastName: FirstLast[2],
		LoginUser:     LoginUser[2],
	},
}

var Contacts = []types.NewContact{
	{
		FirstLastName: FirstLast[0],
		Phone:         "+11234567890",
	}, {
		FirstLastName: FirstLast[1],
		Phone:         "+11234567891",
	}, {
		FirstLastName: FirstLast[2],
		Phone:         "+11234567892",
	},
}
