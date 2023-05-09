package models

import (
	"errors"
	"unicode"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BesforeSave() error {
	if !u.isValid() {
		return errors.New("invalid user info")
	}
	return nil
}

// validates the User model before we save it. conditions:
//
//	the username cannot contain any whitesapce
//	the password must be at least 6 characters, no spaces
func (u *User) isValid() bool {
	for _, rune := range u.Username {
		if unicode.IsSpace(rune) {
			return false
		}
	}

	if len(u.Password) < 6 {
		return false
	}

	for _, rune := range u.Password {
		if unicode.IsSpace(rune) {
			return false
		}
	}

	return true
}
