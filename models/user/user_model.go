package user

import (
	"time"
)

type User struct {
	Id        uint      `jsonapi:"primary,users"`
	Name      string    `jsonapi:"attr, name"`
	LastName  string    `jsonapi:"attr, last_name"`
	Email     string    `jsonapi:"attr, email"`
	CreatedAt time.Time `jsonapi:"attr, created_at"`
	UpdatedAt time.Time `jsonapi:"attr, updated_at"`
}
