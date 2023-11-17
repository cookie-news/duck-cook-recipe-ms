package entity

type CommentRecipe struct {
	Message   string `json:"message"`
	IdComment string `json:"id,omitempty"`
	UserName  string `json:"userName"`
	IdRecipe  string `json:"idRecipe"`
	IdUser    string `json:"idUser"`
}
