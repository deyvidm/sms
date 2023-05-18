package models

import (
	"github.com/deyvidm/sms-backend/types"
	"gorm.io/gorm"
)

const (
	InviteStatus_Sending   = "sending"
	InviteStatus_Invited   = "invited"
	InviteStatus_Accepted  = "accepted"
	InviteStatus_Declined  = "declined"
	InviteStatus_Waitlist  = "waitlist"
	InviteStatus_Uninvited = "uninvited"
)

type Invite struct {
	gorm.Model
	ContactID uint
	Contact   Contact `gorm:"foreignKey:ContactID"`
	EventID   uint
	Event     Event  `gorm:"foreignKey:EventID"`
	Status    string `gorm:"type:text"`
	Paid      bool
}

func UpdateInvite(invite types.UpdateInvite) error {
	return nil
}
