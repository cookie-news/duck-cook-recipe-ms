package entity

import "time"

type CommentRecipe struct {
	Message   string    `json:"message"`
	IdComment string    `json:"id,omitempty"`
	User      User      `json:"user,omitempty"`
	IdRecipe  string    `json:"idRecipe"`
	IdUser    string    `json:"idUser"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
