package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Phone     string `json:"phone"` // newly-registeerd users gain a blank Contact which we map to their outgoing messages
	Owner     uint
}

// used for returning cleaner data structs
// otherwise API consumers also get createdAt, updatedAt, ID, etc.
type APIContact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	ID        string `json:"id"`
}

func (u *User) AllContacts() ([]APIContact, error) {
	contacts := []APIContact{}
	err := DB.Model(u).Association("Contacts").Find(&contacts)
	return contacts, err
}

func (u *User) SaveContact(c Contact) (Contact, error) {
	err := DB.Model(u).Association("Contacts").Append([]Contact{c})
	if err != nil {
		return Contact{}, err
	}
	return c, nil
}
