package models

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	Id        int64     `orm:"auto"`
	Username  string    `orm:"size(100);unique"`
	Email     string    `orm:"size(200)"`
	IsAdmin   bool      `orm:"default(false)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
