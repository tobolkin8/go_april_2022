package auth

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginInfo struct {
	Id                uint   `jsonapi:"primary,id"`
	Username          string `jsonapi:"attr, username"`
	UserId            uint   `jsonapi:"attr, user_id"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `jsonapi:"attr, encrypted_password"`

	CreatedAt   time.Time `jsonapi:"attr, created_at"`
	UpdatedAt   time.Time `jsonapi:"attr, updated_at"`
	LastLoginAt time.Time `jsonapi:"attr, last_login_at"`
}

func (u *LoginInfo) SetPassword(p string) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = ""
	u.EncryptedPassword = string(h)
}

func (u *LoginInfo) ComparePassword(p string) bool {
	var (
		encryptedPassword = []byte(u.EncryptedPassword)
		password          = []byte(p)
	)

	if err := bcrypt.CompareHashAndPassword(encryptedPassword, password); err != nil {
		return false
	}
	return true
}
func (u *LoginInfo) Validate() []error {
	var errs []error

	return errs
}
func (u *LoginInfo) BeforeSave() error {
	u.SetPassword(u.Password)
	return nil
}
