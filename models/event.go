package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/deyvidm/sms-backend/types"
	"gorm.io/gorm"
)

type EventStatus string

const (
	EventStatus_Upcoming  EventStatus = "upcoming"
	EventStatus_Active    EventStatus = "active"
	EventStatus_Completed EventStatus = "completed"
	EventStatus_Cancelled EventStatus = "cancelled"
)

func (e *EventStatus) Scan(value interface{}) error {
	*e = EventStatus(value.([]byte))
	return nil
}

func (e EventStatus) Value() (driver.Value, error) {
	return string(e), nil
}

type Event struct {
	gorm.Model
	Organizer      uint
	Title          string
	InvitationBody string
	TargetCapacity int
	StartDate      *time.Time `gorm:"type:datetime"`
	EndDate        *time.Time `gorm:"type:datetime"`
	InviteDate     *time.Time `gorm:"type:datetime"`
	Status         EventStatus
}

type EventAPI struct {
	Title string `json:"title"`
}

func fetchContacts(owner *User, contactIDs []string) ([]Contact, error) {
	// get all contacts where id in {} and owner=user
	var contacts []Contact
	DB.Where("id IN ? AND owner = ? ", contactIDs, owner.ID).Find(&contacts)
	if len(contactIDs) != len(contacts) {
		return nil, fmt.Errorf("missing contacts: searched for %d but only found %d", len(contactIDs), len(contacts))
	}

	return contacts, nil
}

func (u *User) OrganizeEvent(eventInput types.NewEvent) error {
	contacts, err := fetchContacts(u, eventInput.Contacts)
	if err != nil {
		return err
	}

	transErr := DB.Transaction(func(tx *gorm.DB) error {
		event := Event{
			Title:          eventInput.Title,
			InvitationBody: eventInput.Invitebody,
		}
		if err := tx.Model(u).Association("Events").Append([]Event{event}); err != nil {
			return err
		}

		var invites []Invite
		for _, contact := range contacts {
			invites = append(invites, Invite{
				Contact: contact,
				Event:   event,
			})
		}

		if err = tx.Create(invites).Error; err != nil {
			return err
		}
		return nil
	}).Error()

	if len(transErr) != 0 {
		return fmt.Errorf(transErr)
	}

	for _, c := range contacts {
		// enqueue newMessage task
	}

	return nil
}

func (u *User) SaveEvent(event Event) error {
	return DB.Model(u).Association("Events").Append([]Event{event})
}

func (u *User) AllEvents() (events []EventAPI, err error) {
	err = DB.Model(u).Association("Events").Find(&events)
	return
}
