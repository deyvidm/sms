package models

import (
	"errors"

	"github.com/deyvidm/sms-backend/auth"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username  string `gorm:"size:255;not null;unique"`
	Password  string `gorm:"size:255;not null;"`
	ContactID string
	Contact   Contact   `gorm:"foreignKey:ContactID"`
	Contacts  []Contact `gorm:"foreignKey:Owner"`
	Events    []Event   `gorm:"foreignKey:OrganizerID"`
}

type APIUser struct {
	Username string    `json:"username"`
	Contacts []Contact `json:"contacts"`
	Events   []Event   `json:"events"`
}

func (u *User) ToAPIUser() APIUser {
	return APIUser{
		Username: u.Username,
		Contacts: u.Contacts,
		Events:   u.Events,
	}
}

func GetUserByID(uid string) (User, error) {
	u := User{}
	if err := DB.Preload("Contact").Preload("Contacts").Preload("Events").Where("id = ?", uid).First(&u).Error; err != nil {
		return User{}, errors.New("User not found")
	}
	return u, nil
}

func LoginUser(username string, password string) (user APIUser, token string, err error) {
	u := User{}
	if err = DB.Model(u).Where("username = ?", username).Take(&u).Error; err != nil {
		return
	}
	if err = VerifyPassword(password, u.Password); err != nil {
		return
	}
	token, err = auth.GenerateToken(u.ID)
	return u.ToAPIUser(), token, err
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) RegisterUser() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)

	err = DB.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
