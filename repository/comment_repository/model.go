package comment_repository

import (
	"duck-cook-recipe/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IdUser    primitive.ObjectID `bson:"idUser"`
	IdRecipe  primitive.ObjectID `bson:"idRecipe"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Message   string             `bson:"message"`
}

func (comment Comment) ToEntityComment() entity.CommentRecipe {
	return entity.CommentRecipe{
		IdComment: comment.ID.Hex(),
		IdRecipe:  comment.IdRecipe.Hex(),
		IdUser:    comment.IdUser.Hex(),
		Message:   comment.Message,
		CreatedAt:  comment.CreatedAt,
	}
}

func (Comment) FromEntity(comment entity.CommentRecipe) Comment {
	id, _ := primitive.ObjectIDFromHex(comment.IdComment)
	idUser, _ := primitive.ObjectIDFromHex(comment.IdUser)
	idRecipe, _ := primitive.ObjectIDFromHex(comment.IdRecipe)
	return Comment{
		ID:       id,
		IdUser:   idUser,
		IdRecipe: idRecipe,
		Message:  comment.Message,
		CreatedAt: comment.CreatedAt,
	}
}
