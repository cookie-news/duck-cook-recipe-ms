package entity

type CommentRecipe struct {
	Message   string `json:"message"`
	IdComment string `json:"id,omitempty"`
	User      User   `json:"user"`
	IdRecipe  string `json:"idRecipe"`
	IdUser    string `json:"idUser"`
}
