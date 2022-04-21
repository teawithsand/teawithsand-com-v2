package user

import "time"

type RegisterMailData struct {
	Username  string
	Token     string
	Email     string
	CreatedAt time.Time
}

type InitPasswordResetMailData struct {
	PublicName string
	Token      string
	Email      string
	CreatedAt  time.Time
}
