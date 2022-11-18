package repository

import (
	models "go-postgres/model"
)

type UserRepo interface {
	Select() (users []models.User)
	Create(models.User) error
	Delete(delId int) error
	Update(a *int, users *models.User) error
}
