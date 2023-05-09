package models

import (
	"errors"
	"unicode"

	"github.com/deyvidm/sms-backend/auth"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
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
	u.BesforeSave()
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

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
