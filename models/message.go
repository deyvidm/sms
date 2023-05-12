package models

type Message struct {
	From Contact `gorm:"not null;"`
	To   Contact `gorm:"not null;"`
	Body string
}
