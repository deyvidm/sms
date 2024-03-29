package types

// this file defines Datatypes used by the Gin Backend API as well as our testing files
// NOTE: they are similar to, but separate from, the Models used by GORM and SQLite 3 -- those can be found in /models/

const StatusFailed = "failed"
const StatusSuccess = "success"

type LoginUser struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=255"` // 3 is a holy number
	Password string `json:"password" binding:"required,alphanum,min=6,max=255"` // min 6 for brcypt hash
}

type NewUser struct {
	LoginUser     // every new user (i.e. event organizer) needs valid login info
	FirstLastName // they also receive a Contact entry, which we use to tag the messages they send
}

// an "internal" struct to reduce copy-pasted/repeated code"
type FirstLastName struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=3,max=50"`
	LastName  string `json:"last_name" binding:"required,alpha,min=3,max=50"`
}

type NewContact struct {
	FirstLastName
	Phone string `json:"phone" binding:"required,e164"` // e164 is the standard +11234567890
}

type NewEvent struct {
	Title      string   `json:"title" binding:"required,max=255"`
	Invitebody string   `json:"invite_body" binding:"required"`
	Contacts   []string `json:"contacts" biding:"required"`
	//TODO start/end/invite Dates
	// capacity
}

type NewMessage struct {
	Content string `json:"content" binding:"required"`
	To      string `json:"to" bidnding:"required"` // this will map to a specific Contact.ID
}

// pointers because we want to allow asynq to http.PUT Status without forcing defaults on Paid
type UpdateInvite struct {
	Status *string `json:"status,omitempty" binding:"alpha"`
	Paid   *bool   `json:"paid,omitempty" binding:"boolean"`
}

// This comes in Straight from SNS (AWS)
type SNSEvent struct {
	Type             string `json:"Type"`
	Message          string `json:"Message"` // this is all stringified+escaped JSON => SNSRawMessage
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}
