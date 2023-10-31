package models

import (
	"errors"
	"fmt"

	"github.com/deyvidm/sms/cmd/web-server/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	Username string       `json:"username"`
	Contacts []APIContact `json:"contacts"`
	Events   []APIEvent   `json:"events"`
}

func (u *User) ToAPIUser() APIUser {
	fmt.Println(len(u.Contacts))
	return APIUser{
		Username: u.Username,
		Contacts: Contacts(u.Contacts).ToAPI(),
		Events:   Events(u.Events).ToAPI(),
	}
}

func userPreload() (tx *gorm.DB) {
	return DB.Preload("Contacts").Preload("Events")
}

func GetUserByID(uid string) (User, error) {
	u := User{}
	if err := userPreload().Where("id = ?", uid).First(&u).Error; err != nil {
		return User{}, errors.New("User not found")
	}
	return u, nil
}

func LoginUser(username string, password string) (user APIUser, token string, err error) {
	u := User{}
	if err = DB.Preload("Contact").Preload("Contacts").Preload("Events").Where("username = ?", username).First(&u).Error; err != nil {
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
