package repository

import "duck-cook-recipe/entity"

type UserRepository interface {
	GetUserById(idUser, token string) (user entity.User, err error)
}
