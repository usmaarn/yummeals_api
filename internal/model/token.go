package model

import "time"

type Token struct {
	ID                    int
	UserID                int
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
	UserAgent             string
	IpAddress             string
}
