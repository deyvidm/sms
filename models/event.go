package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title     string
	Organizer uint
}

type EventAPI struct {
	Title string `json:"title"`
}

func (u *User) SaveEvent(event Event) error {
	return DB.Model(u).Association("Events").Append([]Event{event})
}

func (u *User) AllEvents() (events []EventAPI, err error) {
	err = DB.Model(u).Association("Events").Find(&events)
	return
}
