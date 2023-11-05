package like_repository

import (
	"duck-cook-recipe/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IdUser    primitive.ObjectID `bson:"idUser"`
	IdRecipe  primitive.ObjectID `bson:"idRecipe"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (like Like) ToEntityLike() entity.LikeRecipe {
	return entity.LikeRecipe{
		IdLike: like.ID.Hex(),
		Recipe: entity.Recipe{
			Id:     like.IdRecipe.Hex(),
			IdUser: like.IdUser.Hex(),
		},
	}
}

func (Like) FromEntity(like entity.LikeRecipe) Like {
	id, _ := primitive.ObjectIDFromHex(like.IdLike)
	idUser, _ := primitive.ObjectIDFromHex(like.IdUser)
	idRecipe, _ := primitive.ObjectIDFromHex(like.Id)
	return Like{
		ID:       id,
		IdUser:   idUser,
		IdRecipe: idRecipe,
	}
}
