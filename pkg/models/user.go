package models

import "time"

type UserModel struct {
	Id        string        `json:"id"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	CreatedAt time.Duration `json:"created_at"`
	UpdatedAt time.Duration `json:"updated_at"`
}
