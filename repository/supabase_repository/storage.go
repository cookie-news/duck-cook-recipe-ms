package supabase_repository

import (
	"duck-cook-recipe/api/repository"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type storageImpl struct {
	client storage_go.Client
}

func (c *storageImpl) UploadImage(images []*multipart.FileHeader, id string) ([]string, error) {
	var urls []string
	var updateError error
	bucketname := os.Getenv("SUPABASE_BUCKET_ID")

	for _, img := range images {
		file, err := img.Open()
		contentType := img.Header.Get("Content-Type")
		imageName := uuid.NewString()

		if err != nil {
			updateError = err
		}

		result, err := c.client.UploadFile(bucketname, id+"/"+imageName, file, storage_go.FileOptions{ContentType: &contentType})

		if err != nil {
			updateError = err
		}

		url := c.GetPublicUrl(result.Key)

		urls = append(urls, url)
	}

	return urls, updateError
}

func (c *storageImpl) GetPublicUrl(filename string) string {
	return "https://dcqgxwhjxkkignjziqvf.supabase.co/storage/v1/object/public/" + filename
}

func New(client storage_go.Client) repository.RecipeStorage {
	return &storageImpl{client}
}
