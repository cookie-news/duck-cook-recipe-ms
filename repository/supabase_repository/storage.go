package supabase_repository

import (
	"duck-cook-recipe/api/repository"
	"log"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type storageImpl struct {
	client     storage_go.Client
	bucketname string
}

func (c *storageImpl) UploadImage(images []*multipart.FileHeader, id string) ([]string, error) {
	var urls []string
	var updateError error

	for _, img := range images {
		file, err := img.Open()
		contentType := img.Header.Get("Content-Type")
		imageName := uuid.NewString()

		if err != nil {
			updateError = err
		}

		result, err := c.client.UploadFile(c.bucketname, id+"/"+imageName, file, storage_go.FileOptions{ContentType: &contentType})

		if err != nil {
			updateError = err
		}

		url := c.GetPublicUrl(result.Key, "", "")

		urls = append(urls, url)
	}

	return urls, updateError
}

func (c *storageImpl) GetPublicUrl(filename string, bucketname string, folderName string) string {
	baseUrl := "https://dcqgxwhjxkkignjziqvf.supabase.co/storage/v1/object/public/"
	if bucketname != "" {
		return baseUrl + bucketname + "/" + folderName + "/" + filename
	}
	return baseUrl + filename
}

func (c *storageImpl) ListFiles(folderName string) (files []string, err error) {
	filesStorage, err := c.client.ListFiles(c.bucketname, folderName, storage_go.FileSearchOptions{})

	if err != nil {
		log.Println("[SUPABASE]", err.Error())
		return files, err
	}

	for _, image := range filesStorage {
		files = append(files, c.GetPublicUrl(image.Name, c.bucketname, folderName))
	}

	return
}

func New(client storage_go.Client) repository.RecipeStorageRepository {
	bucketname := os.Getenv("SUPABASE_BUCKET_ID")
	return &storageImpl{client, bucketname}
}
