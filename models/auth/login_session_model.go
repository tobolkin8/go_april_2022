package auth

import (
	"time"
)

type UserSession struct {
	Id          uint   `jsonapi:"primary,login_info"`
	LoginInfoId string `jsonapi:"attr, login_unfo_id"`
	JWT_Token   string `jsonapi:"attr, jwt_token"`

	CreatedAt time.Time `jsonapi:"attr, created_at"`
	UpdatedAt time.Time `jsonapi:"attr, updated_at"`
}
