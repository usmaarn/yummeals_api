package model

import "time"

type User struct {
	ID                    int
	Name                  string
	Email                 string
	PhoneNumber           string
	Password              string
	Status                string
	CreatedAt             time.Time
	EmailVerifiedAt       *time.Time
	PhoneNumberVerifiedAt *time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
}
