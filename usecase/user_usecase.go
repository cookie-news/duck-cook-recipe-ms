package usecase

import (
	"context"
	"duck-cook-recipe/api/repository"
	"duck-cook-recipe/entity"
	"encoding/hex"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/sha3"
)

type UserUseCase interface {
	GetUserById(idUser, token string) (user entity.User, err error)
}

type userUseCaseImpl struct {
	userRepository repository.UserRepository
	redisClient    *redis.Client
}

func CalculateSHA3(input string) string {
	hasher := sha3.New256()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	hashInHex := hex.EncodeToString(hash)
	return hashInHex
}

func (usecase userUseCaseImpl) GetUserById(idUser, token string) (user entity.User, err error) {
	userIdKey := "user:" + CalculateSHA3(idUser)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := usecase.redisClient.HGetAll(ctx, userIdKey).Result()
	if err != redis.Nil &&
		err != nil {
		return
	}

	if result != nil && result["ID"] != "" {
		mapstructure.Decode(result, &user)
		return
	}

	user, err = usecase.userRepository.GetUserById(idUser, token)
	if err != nil {
		return
	}

	var resultMap map[string]interface{}
	mapstructure.Decode(user, &resultMap)
	usecase.redisClient.HMSet(ctx, userIdKey, resultMap).Err()

	err = usecase.redisClient.Expire(ctx, userIdKey, 10*time.Minute).Err()

	return
}

func NewUserUseCase(
	userRepository repository.UserRepository,
	redisClient *redis.Client) UserUseCase {
	return &userUseCaseImpl{
		userRepository,
		redisClient,
	}
}
