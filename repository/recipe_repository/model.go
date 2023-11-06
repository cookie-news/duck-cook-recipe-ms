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
	Ingredients     []Ingredient       `bson:"ingredients"`
}

type Ingredient struct {
	Name     string  `bson:"name"`
	Quantity float64 `bson:"qty"`
	Measure  string  `bson:"measure"`
}

func (ingredient Ingredient) ToEntity() entity.Ingredients {
	return entity.Ingredients{
		Name:    ingredient.Name,
		Qty:     ingredient.Quantity,
		Measure: ingredient.Measure,
	}
}

func (Ingredient) FromEntity(ingredient entity.Ingredients) Ingredient {
	return Ingredient{
		Name:     ingredient.Name,
		Quantity: ingredient.Qty,
		Measure:  ingredient.Measure,
	}
}

func (recipe Recipe) ToEntityRecipe() entity.Recipe {

	var ingredients []entity.Ingredients

	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, ingredient.ToEntity())
	}

	return entity.Recipe{
		Id:              recipe.ID.Hex(),
		IdUser:          recipe.IdUser.Hex(),
		Title:           recipe.Title,
		Description:     recipe.Description,
		PreparationTime: recipe.PreparationTime,
		ImageUrl:        recipe.ImageUrl,
		Ingredients:     ingredients,
	}
}

func (Recipe) FromEntity(recipe entity.Recipe) *Recipe {
	id, _ := primitive.ObjectIDFromHex(recipe.Id)
	idUser, _ := primitive.ObjectIDFromHex(recipe.IdUser)

	var ingredients []Ingredient

	for _, ingredient := range recipe.Ingredients {
		var ingredientModel Ingredient
		ingredients = append(ingredients, ingredientModel.FromEntity(ingredient))
	}

	return &Recipe{
		ID:              id,
		IdUser:          idUser,
		Title:           recipe.Title,
		Description:     recipe.Description,
		PreparationTime: recipe.PreparationTime,
		ImageUrl:        recipe.ImageUrl,
		Ingredients:     ingredients,
	}
}
