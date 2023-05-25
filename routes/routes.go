package routes

import (
	"github.com/deyvidm/sms-backend/controllers"
	"github.com/gin-gonic/gin"
)

const APIPrefix = "/api"

const PublicPrefix = APIPrefix

const Ping = "/ping"

const UserLogin = "/users/login"
const UserRegister = "/users/register"
const CurrentUser = "/user"

const SNSEvent = "/respond"

func AssignPublicRoutes(router *gin.RouterGroup) {
	router.GET(Ping, controllers.Pong)
	router.GET(CurrentUser, controllers.CurrentUser)
	router.POST(UserRegister, controllers.Register)
	router.POST(UserLogin, controllers.Login)
	router.POST(SNSEvent, controllers.ReceiveSNS)
}

const PrivatePrefix = APIPrefix

const NewContact = "/contacts/new"
const AllContacts = "/contacts"

const NewEvent = "/events/new"
const AllEvents = "/events"
const EventDetails = "/events/:id"

const NewMessage = "/messages/new"

func AssignPrivateRoutes(router *gin.RouterGroup) {
	router.POST(NewContact, controllers.NewContact)
	router.GET(AllContacts, controllers.AllContacts)

	router.POST(NewEvent, controllers.NewEvent)
	router.GET(AllEvents, controllers.AllEvents)
	router.GET(EventDetails, controllers.EventDetails)

	router.POST(NewMessage, controllers.NewMessage)
}

const InternalPrefix = APIPrefix + "/internal"

const UpdateInvite = "/invite/:id"

func AssignInternalRoutes(router *gin.RouterGroup) {
	router.PUT(UpdateInvite, controllers.UpdateInvite)
}
