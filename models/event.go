package models

import (
	"fmt"
	"log"
	"time"

	"github.com/deyvidm/sms-asynq/tasks"
	"github.com/deyvidm/sms-backend/types"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

const (
	EventStatus_Upcoming  = "upcoming"
	EventStatus_Active    = "active"
	EventStatus_Completed = "completed"
	EventStatus_Cancelled = "cancelled"
)

type Event struct {
	BaseModel
	OrganizerID    string
	Title          string
	InvitationBody string
	TargetCapacity int
	StartDate      *time.Time `gorm:"type:datetime"`
	EndDate        *time.Time `gorm:"type:datetime"`
	InviteDate     *time.Time `gorm:"type:datetime"`
	Status         string     `gorm:"type:text"`
}

type APIEvent struct {
	ID             string     `json:"id"`
	Title          string     `json:"title"`
	TargetCapacity int        `json:"capacity"`
	StartDate      *time.Time `json:"start_date"`
	EndDate        *time.Time `json:"type:end_date"`
}

func (e *Event) ToAPIEvent() APIEvent {
	return APIEvent{
		ID:             e.ID,
		Title:          e.Title,
		TargetCapacity: e.TargetCapacity,
		StartDate:      e.StartDate,
		EndDate:        e.EndDate,
	}
}

type Events []Event

func (events Events) toAPIEvent() []APIEvent {
	var ret []APIEvent
	for _, c := range events {
		ret = append(ret, c.ToAPIEvent())
	}
	return ret
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

func EventFromInput(e types.NewEvent) Event {
	return Event{
		Title:          e.Title,
		InvitationBody: e.Invitebody,
		Status:         EventStatus_Active,
	}
}

func (u *User) OrganizeEvent(eventInput types.NewEvent) error {
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	contacts, err := fetchContacts(u, eventInput.Contacts)
	if err != nil {
		return err
	}

	return DB.Transaction(func(tx *gorm.DB) error {
		event := EventFromInput(eventInput)
		event.OrganizerID = u.ID

		var invites []Invite
		for _, contact := range contacts {
			invites = append(invites, Invite{
				ContactID: contact.ID,
				Event:     event,
				Status:    InviteStatus_Sending,
			})
		}
		if err := tx.Create(&invites).Error; err != nil {
			return err
		}

		for _, inv := range invites {
			for _, contact := range contacts {
				if contact.ID != inv.ContactID {
					continue
				}
				t, err := tasks.NewInviteTask(inv.ID, contact.Phone, eventInput.Invitebody)
				if err != nil {
					return err
				}

				taskInfo, err := asynqClient.Enqueue(t)
				if err != nil {
					return err
				}

				log.Printf("enqued task %s | inviting %s to %s...", taskInfo.ID, inv.Contact.Phone, eventInput.Title)
			}
		}
		return nil
	})
}

func (u *User) AllEvents() (events []APIEvent, err error) {
	err = DB.Model(u).Association("Events").Find(&events)
	return
}
