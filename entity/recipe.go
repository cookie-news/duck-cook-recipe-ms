package entity

import "mime/multipart"

type Recipe struct {
	Id                string                  `form:"id" json:"id"`
	IdUser            string                  `form:"idUser" json:"idUser" binding:"required"`
	Title             string                  `form:"title" json:"title" binding:"required"`
	Description       string                  `form:"description" json:"description" binding:"required"`
	PreparationMethod string                  `form:"preparationMethod" binding:"required" json:"preparationMethod"`
	PreparationTime   int                     `form:"preparationTime" json:"preparationTime" example:"600" format:"int64" binding:"required"`
	Images            []*multipart.FileHeader `form:"images" example:"arquivos de imagens" json:"images" format:"blob" binding:"required" swaggerignore:"true"`
	Ingredients       []string                `form:"ingredients" json:"ingredients"`
}

type RecipeResponse struct {
	Recipe
	Images []string `json:"images"`
}

type RecipeCountLikeManyComments struct {
	RecipeResponse
	CountLikes    int `json:"countLikes"`
	CountComments int `json:"countComments"`
}
