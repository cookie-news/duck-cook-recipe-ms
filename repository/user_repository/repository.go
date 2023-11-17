package user_repository

import (
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

type UserRepositoryImpl struct {
	clientHttp *resty.Client
}

func (repository UserRepositoryImpl) GetUserById(idUser, token string) (user entity.User, err error) {
	resp, err := repository.clientHttp.R().
		SetHeader("authorization", token).
		Get("/v1/customer/_id/" + idUser)

	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Body(), &user)
	return
}

func NewUserRepositoryImpl() repository.UserRepository {
	client := resty.New()
	client.BaseURL = os.Getenv("URL_USER")
	return &UserRepositoryImpl{
		client,
	}
}
