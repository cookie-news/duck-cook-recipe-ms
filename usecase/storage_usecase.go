package usecase

import (
	"context"
	"duck-cook-recipe/api/repository"
	"fmt"
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

	usecase.saveStringArray(folderIdKey, files)

	return
}

func (usecase *storageUseCaseImpl) saveStringArray(key string, values []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	data := make(map[string]interface{})
	for i, value := range values {
		data[fmt.Sprintf("%s:%d", key, i)] = value
	}
	_, err := usecase.redisClient.MSet(ctx, data).Result()
	return err
}

func (usecase *storageUseCaseImpl) getStringArray(key string) ([]string, error) {
	ctx := context.Background()

	keys, err := usecase.redisClient.Keys(ctx, fmt.Sprintf("%s:*", key)).Result()
	if err != nil {
		return nil, err
	}

	result, err := usecase.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	var values []string
	for _, v := range result {
		if v != nil {
			values = append(values, v.(string))
		}
	}

	return values, nil
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
