package repository

import (
	"mime/multipart"
)

type RecipeStorageRepository interface {
	UploadImage(image []*multipart.FileHeader, id string) ([]string, error)
	GetPublicUrl(filename string, bucketname string, foldername string) string
	ListFiles(folderName string) (files []string, err error)
}
