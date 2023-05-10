package models

import (
	"errors"

	"github.com/deyvidm/sms-backend/auth"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Password string    `gorm:"size:255;not null;" json:"password"`
	Contacts []Contact `gorm:"foreignKey:Owner"`
}

func GetUserByID(uid uint) (User, error) {
	u := User{}
	if err := DB.First(&u, uid).Error; err != nil {
		return User{}, errors.New("User not found")
	}
	return u, nil
}

func LoginUser(username string, password string) (token string, err error) {
	u := User{}
	if err = DB.Model(u).Where("username = ?", username).Take(&u).Error; err != nil {
		return
	}
	if err = VerifyPassword(password, u.Password); err != nil {
		return
	}
	return auth.GenerateToken(u.ID)
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) SaveUser() (*User, error) {
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
