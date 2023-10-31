package models

import (
	"fmt"

	"github.com/deyvidm/sms/web-server/types"
)

type Message struct {
	BaseModel
	SenderID    string
	RecipientID string
	From        Contact `gorm:"not null;foreignKey:SenderID"`
	To          Contact `gorm:"not null;foreignKey:RecipientID"`
	Body        string
}

func (u *User) SaveMessage(input types.NewMessage) error {
	var recipient Contact
	DB.Where("ID = ? AND owner = ? ", input.To, u.ID).First(&recipient)
	if recipient == (Contact{}) {
		return fmt.Errorf("missing contact")
	}
	msg := Message{
		SenderID:    u.ContactID,
		RecipientID: recipient.ID,
		Body:        input.Content,
	}
	return DB.Create(&msg).Error
}
