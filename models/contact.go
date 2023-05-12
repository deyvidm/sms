package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Phone     string `json:"phone"` // newly-registered users gain a blank Contact which we map to their outgoing messages
	Owner     uint
}

func (c *Contact) toAPIContact() APIContact {
	return APIContact{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
		ID:        c.ID,
	}
}

// used for returning cleaner data structs
// otherwise API consumers also get createdAt, updatedAt, ID, etc.
type APIContact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	ID        uint   `json:"id"`
}

func (u *User) AllContacts() ([]APIContact, error) {
	contacts := []APIContact{}
	err := DB.Model(u).Association("Contacts").Find(&contacts)
	return contacts, err
}

func (u *User) SaveContact(c Contact) (APIContact, error) {
	var contact Contact
	DB.Where("phone = ? AND owner = ? ", c.Phone, u.ID).First(&contact)
	if contact != (Contact{}) {
		return contact.toAPIContact(), nil
	}
	err := DB.Model(u).Association("Contacts").Append([]Contact{c})
	if err != nil {
		return APIContact{}, err
	}
	DB.Where("phone = ? AND owner = ? ", c.Phone, u.ID).First(&contact)
	return contact.toAPIContact(), nil
}
