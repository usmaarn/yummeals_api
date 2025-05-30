package model

import "time"

type Token struct {
	ID                    int64     `db:"id" json:"-"`
	UserID                int64     `db:"user_id" json:"-"`
	AccessToken           string    `db:"access_token" json:"access_token"`
	RefreshToken          string    `db:"refresh_token" json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `db:"access_token_expires_at" json:"-"`
	RefreshTokenExpiresAt time.Time `db:"refresh_token_expires_at" json:"-"`
	UserAgent             string    `db:"user_agent" json:"-"`
	IpAddress             string    `db:"ip_address" json:"-"`
}
