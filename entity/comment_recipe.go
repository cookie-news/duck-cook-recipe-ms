package entity

type CommentRecipe struct {
	Message   string `json:"message"`
	IdComment string `json:"id,omitempty"`
	Recipe    `json:"-"`
}
