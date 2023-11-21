package usecase

import (
	"context"
	"duck-cook-recipe/api/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageUseCase interface {
	ListFiles(folderName string) (files []string, err error)
}

type storageUseCaseImpl struct {
	recipeStorageRepository repository.RecipeStorageRepository
	redisClient             *redis.Client
}

func (usecase *storageUseCaseImpl) ListFiles(folderName string) (files []string, err error) {
	folderIdKey := "folder:" + CalculateSHA3(folderName)
	files, err = usecase.getStringArray(folderIdKey)
	if err != redis.Nil &&
		err != nil {
		return
	}
	if len(files) > 0 {
		return
	}

	files, err = usecase.recipeStorageRepository.ListFiles(folderName)
	if err != nil {
		return
	}

	if len(files) > 0 {
		usecase.saveStringArray(folderIdKey, files)
	}
	return
}

func (usecase *storageUseCaseImpl) saveStringArray(key string, values []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := usecase.redisClient.RPush(ctx, key, values).Result()
	if err != nil {
		return err
	}
	ttl := 60 * time.Minute
	_, err = usecase.redisClient.Expire(ctx, key, ttl).Result()
	return err
}

func (usecase *storageUseCaseImpl) getStringArray(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := usecase.redisClient.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewStorageUseCase(
	recipeStorageRepository repository.RecipeStorageRepository,
	redisClient *redis.Client,
) StorageUseCase {
	return &storageUseCaseImpl{
		recipeStorageRepository,
		redisClient,
	}
}
