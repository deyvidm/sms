package routes

import (
	"github.com/deyvidm/sms-backend/controllers"
	"github.com/gin-gonic/gin"
)

const Ping = "/ping"
const UserLogin = "/users/login"
const UserRegister = "/users/register"
const CurrentUser = "/user"

func AssignPublicRoutes(router *gin.RouterGroup) {
	router.GET(Ping, controllers.Pong)
	router.POST(UserRegister, controllers.Register)
	router.POST(UserLogin, controllers.Login)
}

const NewContact = "/contacts/new"
const AllContacts = "/contacts"

const NewEvent = "/events/new"
const AllEvents = "/events"

const NewMessage = "/messages/new"

func AssignPrivateRoutes(router *gin.RouterGroup) {
	router.GET(CurrentUser, controllers.CurrentUser)
	router.POST(NewContact, controllers.NewContact)
	router.GET(AllContacts, controllers.AllContacts)

	router.POST(NewEvent, controllers.NewEvent)
	router.GET(AllEvents, controllers.AllEvents)

	router.POST(NewMessage, controllers.NewMessage)
}

const UpdateInvite = "/invite/:id"

func AssignInternalRoutes(router *gin.RouterGroup) {
	router.PUT(UpdateInvite, controllers.UpdateInvite)
}
