package models

import (
	"database/sql/driver"
	"time"

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

func (u *User) OrganizeEvent(event Event) error {
	// 1. open SQL transaction
	// 2. create Event
	// 3. for each contact
	// 		a. create Invitation
	// 		b. queue NewMessage Event to asynq
	// submit transaction
	return nil
}

func (u *User) SaveEvent(event Event) error {
	return DB.Model(u).Association("Events").Append([]Event{event})
}

func (u *User) AllEvents() (events []EventAPI, err error) {
	err = DB.Model(u).Association("Events").Find(&events)
	return
}
