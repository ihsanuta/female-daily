package model

import "time"

type FetchUserParam struct {
	Page int `json:"page" form:"page"`
}

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email" validate:"required,email"`
	FirstName string     `json:"first_name" validate:"required"`
	LastName  string     `json:"last_name" validate:"required"`
	Avatar    string     `json:"avatar" validate:"required,url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ResponseFetch struct {
	Data []*User `json:"data"`
}

type UserParam struct {
	Email     string `json:"email" form:"email"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Limit     int    `json:"limit" form:"limit"`
	Page      int    `json:"page" form:"page"`
}
