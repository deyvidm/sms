package models

import (
	"github.com/deyvidm/sms/cmd/web-server/types"
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
	Status    string `gorm:"type:text"`
	Paid      bool
}

type APIInvite struct {
	ID      string     `json:"id"`
	Contact APIContact `json:"contact"`
	Status  string     `json:"status"`
	Paid    bool       `json:"paid"`
}

func (i *Invite) ToAPI() APIInvite {
	return APIInvite{
		ID:      i.ID,
		Contact: i.Contact.ToAPI(),
		Status:  i.Status,
		Paid:    i.Paid,
	}
}

func (u *User) GetInviteByID(id string) (Invite, error) {
	var invite Invite
	DB.Where("invites.id = ?", id).First(&invite)
	// Prefill Event and check Organizer ID == user to stop user info leak

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
