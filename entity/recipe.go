package entity

type Recipe struct {
	Id              string `json:"id"`
	IdUser          string `json:"idUser"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	PreparationTime int    `json:"preparationTime" example:"600" format:"int64"`
	ImageUrl        string `json:"imageUrl" example:"url da imagem" format:"string"`
}
