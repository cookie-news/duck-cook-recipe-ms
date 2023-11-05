package recipe_repository

import (
	"duck-cook-recipe/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	IdUser          primitive.ObjectID `bson:"idUser"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
	Title           string             `bson:"title"`
	Description     string             `bson:"description"`
	PreparationTime int                `bson:"preparationTime"`
	ImageUrl        string             `bson:"imageUrl"`
}

func (recipe Recipe) ToEntityRecipe() entity.Recipe {
	return entity.Recipe{
		Id:              recipe.ID.Hex(),
		IdUser:          recipe.IdUser.Hex(),
		Title:           recipe.Title,
		Description:     recipe.Description,
		PreparationTime: recipe.PreparationTime,
		ImageUrl:        recipe.ImageUrl,
	}
}

func (Recipe) FromEntity(recipe entity.Recipe) *Recipe {
	id, _ := primitive.ObjectIDFromHex(recipe.Id)
	idUser, _ := primitive.ObjectIDFromHex(recipe.IdUser)
	return &Recipe{
		ID:              id,
		IdUser:          idUser,
		Title:           recipe.Title,
		Description:     recipe.Description,
		PreparationTime: recipe.PreparationTime,
		ImageUrl:        recipe.ImageUrl,
	}
}
