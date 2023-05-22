package models

type Contact struct {
	BaseModel
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Phone     string `json:"phone"` // newly-registered users gain a blank Contact which we map to their outgoing messages
	Owner     string
}

// used for returning cleaner data structs
// otherwise API consumers also get createdAt, updatedAt, ID, etc.
type APIContact struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (c *Contact) toAPIContact() APIContact {
	return APIContact{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
	}
}

type Contacts []Contact

func (contacts Contacts) toAPIContact() []APIContact {
	var ret []APIContact
	for _, c := range contacts {
		ret = append(ret, c.toAPIContact())
	}
	return ret
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
