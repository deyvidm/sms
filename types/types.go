package types

type InviteStatus string

var InviteStatus_Sending InviteStatus = "sending"
var InviteStatus_Invited InviteStatus = "invited"
var InviteStatus_Accepted InviteStatus = "accepted"
var InviteStatus_Declined InviteStatus = "declined"
var InviteStatus_Waitlist InviteStatus = "waitlist"
var InviteStatus_Uninvited InviteStatus = "uninvited"

func (s InviteStatus) String() string {
	return string(s)
}

type ResponseInfo struct {
	Status          *InviteStatus
	TargetInviteKey *float64
}
