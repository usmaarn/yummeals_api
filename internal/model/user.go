package model

import "time"

type User struct {
	ID                    int64      `db:"id" json:"id"`
	Name                  string     `db:"name" json:"name"`
	Email                 string     `db:"email" json:"email"`
	PhoneNumber           string     `db:"phone_number" json:"phone_number"`
	Password              string     `db:"password" json:"-"`
	Status                int        `db:"status" json:"status"` // e.g., "active", "inactive", "banned"
	CreatedAt             time.Time  `db:"created_at" json:"created_at"`
	EmailVerifiedAt       *time.Time `db:"email_verified_at" json:"email_verified_at,omitempty"`
	PhoneNumberVerifiedAt *time.Time `db:"phone_number_verified_at" json:"phone_number_verified_at,omitempty"`
	UpdatedAt             time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt             *time.Time `db:"deleted_at" json:"deleted_at,omitempty"` // Soft delete field
}
