package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type InviteStatus string

const (
	Sending   InviteStatus = "sending"
	Invited   InviteStatus = "invited"
	Accepted  InviteStatus = "accepted"
	Declined  InviteStatus = "declined"
	Waitlist  InviteStatus = "waitlist"
	Uninvited InviteStatus = "uninvited"
)

func (e *InviteStatus) Scan(value interface{}) error {
	*e = InviteStatus(value.([]byte))
	return nil
}

func (e InviteStatus) Value() (driver.Value, error) {
	return string(e), nil
}

type Invite struct {
	gorm.Model
	ContactID      int
	Contact        Contact `gorm:"foreignKey:ContactID"`
	EventID        int
	Event          Event `gorm:"foreignKey:EventID"`
	Status         InviteStatus
	TargetCapacity int
	Paid           bool
}
