package entity

type CommentRecipe struct {
	Message string `json:"message"`
	Id      string `json:"id,omitempty"`
	Recipe  `json:"-"`
}
