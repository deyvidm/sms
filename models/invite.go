package models

import (
	"github.com/deyvidm/sms-backend/types"
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
	BaseModel
	ContactID string
	Contact   Contact `gorm:"foreignKey:ContactID"`
	EventID   string
	Event     Event  `gorm:"foreignKey:EventID"`
	Status    string `gorm:"type:text"`
	Paid      bool
}

type APIInvite struct {
	ID      string     `json:"id"`
	Contact APIContact `json:"contact"`
	Event   APIEvent   `json:"-"`
	Status  string     `status:"status"`
	Paid    bool       `json:"paid"`
}

func (i *Invite) ToAPI() APIInvite {
	return APIInvite{
		ID:      i.ID,
		Contact: i.Contact.ToAPI(),
		Event:   i.Event.ToAPI(),
		Status:  i.Status,
		Paid:    i.Paid,
	}
}

func GetInvite(id string) (Invite, error) {
	var invite Invite
	DB.Where("id = ?", id).First(&invite)
	// if invite == (Invite{}) {
	// 	return invite, fmt.Errorf("no invite with ID '%s'", id)
	// }
	return invite, nil

}

func (i *Invite) Save(invite types.UpdateInvite) error {
	if invite.Paid != nil {
		i.Paid = *invite.Paid
	}
	if invite.Status != nil {
		i.Status = *invite.Status
	}
	return DB.Save(i).Error
}
