package model

type User struct {
	ID     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Email  string `json:"email,omitempty"`
}
